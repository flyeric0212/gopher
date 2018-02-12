package rpc

import (
	"container/list"
	"errors"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	errPoolClosed = errors.New("rpc: connection pool closed")
	errConnClosed = errors.New("rpc: connection closed")
)

type Pool struct {
	MaxIdle int

	MaxActive int

	Wait        bool
	IdleTimeout time.Duration
	Dial        func() (*Client, error)
	// mu protects fields defined below.
	mu     sync.Mutex
	cond   *sync.Cond
	closed bool
	active int

	// Stack of idleConn with most recently used at the front.
	idle list.List
}

func NewPool(dial func() (*Client, error), maxIdle, maxActive int, idelTimeout time.Duration, wait bool) *Pool {
	return &Pool{Dial: dial, MaxIdle: maxIdle, MaxActive: maxActive, IdleTimeout: idelTimeout, Wait: wait}
}

// conn returns a newly-opened or cached *driverConn
func (pool *Pool) Get() (*PoolClient, error) {
	c, err := pool.get()

	return &PoolClient{pool: pool, client: c}, err
}

func (p *Pool) release() {
	p.active -= 1
	if p.cond != nil {
		p.cond.Signal()
	}
}

// Close releases the resources used by the pool.
func (p *Pool) Close() error {
	p.mu.Lock()
	idle := p.idle
	p.idle.Init()
	p.closed = true
	p.active -= idle.Len()
	if p.cond != nil {
		p.cond.Broadcast()
	}
	p.mu.Unlock()
	for e := idle.Front(); e != nil; e = e.Next() {
		e.Value.(*Client).Close()
	}
	return nil
}

func (p *Pool) get() (*Client, error) {
	p.mu.Lock()

	// Check for pool closed before dialing a new connection.
	if p.closed {
		p.mu.Unlock()
		return nil, errors.New("rpc pool: get on closed pool")
	}

	for {

		if p.closed {
			p.mu.Unlock()
			return nil, errors.New("rpc pool: get on closed pool")
		}

		// Get idle connection.
		for i, n := 0, p.idle.Len(); i < n; i++ {
			e := p.idle.Front()
			if e == nil {
				break
			}
			ic := e.Value.(*Client)
			p.idle.Remove(e)
			p.mu.Unlock()
			return ic, nil
		}

		// Dial new connection if under limit.
		if p.MaxActive == 0 || p.active < p.MaxActive {
			dial := p.Dial
			p.active += 1
			p.mu.Unlock()
			c, err := dial()
			if err != nil {
				p.mu.Lock()
				p.release()
				p.mu.Unlock()
				c = nil
			}
			return c, err
		}

		if !p.Wait {
			p.mu.Unlock()
			return nil, errors.New("pool: connection pool exhausted")
		}

		if p.cond == nil {
			p.cond = sync.NewCond(&p.mu)
		}
		// fmt.Println("rpc conn pool, time: ", time.Now().Local())
		p.cond.Wait()
	}

	return nil, nil
}

func (p *Pool) put(c *Client, err error) error {

	p.mu.Lock()

	if p.closed {
		p.mu.Unlock()
		return nil
	}

	if c == nil {
		p.release()
		if nil != p.cond {
			p.cond.Signal()
		}
		p.mu.Unlock()
		return nil
	}

	if (nil != err && (isNetError(err)) || isTimeoutError(err)) || c.closing || c.shutdown || p.idle.Len() >= p.MaxIdle {
		c.Close()
		p.release()
		if nil != p.cond {
			p.cond.Signal()
		}
		p.mu.Unlock()
		return nil
	} else {
		p.idle.PushFront(c)
	}

	if nil != p.cond {
		p.cond.Signal()
	}
	p.mu.Unlock()
	return nil
}

type PoolClient struct {
	pool   *Pool
	client *Client
	err    error
}

func (pc *PoolClient) Close() {
	pc.pool.put(pc.client, pc.err)
}

func (pc *PoolClient) CallTimeout(serviceMethod string, args interface{}, reply interface{}, retry int, duration time.Duration) error {
	var err error
	for i := 0; i < retry; i++ {
		err = pc.client.CallTimeout(serviceMethod, args, reply, duration)
		if nil == err {
			break
		} else if isNetError(err) {
			log.Println("net err , reconnect")
			//pc.client.Close()
			dial := pc.pool.Dial
			pc.client, err = dial()
			if nil != err {
				err = ErrShutdown
				break
			}
		} else if isTimeoutError(err) {
			//兼容阿里云网络问题
			//pc.client.Close()
			//log.Println("timeout , reconnect")
			//dial := pc.pool.Dial
			//pc.client, err = dial()
			//if nil != err {
			//err = ErrShutdown
			//超时不重试 直接失败
			break
			//}
		} else {
			break
		}
	}
	pc.err = err
	return err
}

func isTimeoutError(err error) bool {
	if err == nil {
		return false
	}

	operr, ok := err.(*net.OpError)
	if ok && operr.Timeout() {
		return true
	} else if err == errTimeout {
		return true
	} else if strings.Contains(err.Error(), errTimeout.Error()) {
		return true
	}
	return false
}

func isNetError(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*net.OpError)
	if ok {
		return true
	}

	if err == io.EOF || err == io.ErrUnexpectedEOF {
		return true
	}

	if strings.Contains(err.Error(), io.ErrUnexpectedEOF.Error()) {
		return true
	}

	if strings.Contains(err.Error(), "connection is shut down") {
		return true
	}

	if strings.Contains(err.Error(), "connection reset by peer") {
		return true
	}

	if strings.Contains(err.Error(), "invalid character") && strings.Contains(err.Error(), "looking for beginning of value") {
		return true
	}

	if strings.Contains(err.Error(), "use of closed network connection") {
		return true
	}

	return false
}
