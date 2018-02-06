package rpc

import (
	"container/heap"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type Log interface {
	Error(format string, args ...interface{})
	Info(format string, args ...interface{})
	Notice(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Debug(format string, args ...interface{})
}

func (c *RpcClusterClient) Errorf(format string, args ...interface{}) {
	if nil != c.logger {
		c.logger.Error(format, args...)
	}
}

func (c *RpcClusterClient) Debugf(format string, args ...interface{}) {
	if nil != c.logger {
		//c.logger.Debug(format, args...)
	}
}

func (c *RpcClusterClient) Warningf(format string, args ...interface{}) {
	if nil != c.logger {
		c.logger.Warning(format, args...)
	}
}

const errWeight uint64 = 10
const minHeapSize = 1

type weightClientStats struct {
	UseCount   int64 `json:"use_count"`
	ErrorCount int64 `json:"error_count"`
}

type poolWeightClient struct {
	pool          *Pool
	endpoint      string // "127.0.0.1"
	port          string // "127.0.0.1"
	index         int
	weight        uint64
	errcnt        int
	clusterClient *RpcClusterClient
	stats         weightClientStats
}

func (client *poolWeightClient) Connect() (*Client, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", client.endpoint, client.port), client.clusterClient.timeout)
	if err != nil {
		client.clusterClient.Errorf("get %s rpc client error :%v", client.clusterClient.addr, err)
		return nil, err
	}

	rpc_client := client.clusterClient.getClient(conn)

	return rpc_client, nil
}

func NewRpcClusterClient(getClient func(conn io.ReadWriteCloser) *Client, addr string, logger Log, timeout time.Duration, retry int) *RpcClusterClient {
	c := &RpcClusterClient{
		getClient: getClient,
		timeout:   timeout,
		addr:      addr,
		logger:    logger,
		retry:     retry,
	}

	c.updateClientAddr()
	go func(c *RpcClusterClient) {
		timer := time.NewTicker(30 * time.Second)
		for {
			select {
			case <-timer.C:
				c.updateClientAddr()
			}
		}
	}(c)

	return c
}

type RpcClusterClient struct {
	logger    Log
	getClient func(conn io.ReadWriteCloser) *Client
	timeout   time.Duration
	retry     int
	addr      string
	sync.RWMutex
	endpoints []string
	clients   []*poolWeightClient
}

func (c *RpcClusterClient) Len() int {
	return len(c.clients)
}

func (c *RpcClusterClient) Swap(i, j int) {
	c.clients[i], c.clients[j] = c.clients[j], c.clients[i]
	c.clients[i].index = i
	c.clients[j].index = j
}

func (c *RpcClusterClient) Less(i, j int) bool {
	return c.clients[i].weight < c.clients[j].weight
}

func (c *RpcClusterClient) Pop() (client interface{}) {
	c.clients, client = c.clients[:c.Len()-1], c.clients[c.Len()-1]
	return
}

func (c *RpcClusterClient) Push(client interface{}) {
	weightClient := client.(*poolWeightClient)
	weightClient.index = c.Len()
	c.clients = append(c.clients, weightClient)
}

func (c *RpcClusterClient) exist(addr string) bool {
	c.RLock()
	for _, cli := range c.clients {
		if cli.endpoint == addr {
			c.RUnlock()
			return true
		}
	}
	c.RUnlock()
	return false
}

func (c *RpcClusterClient) add(addr string, weightClient *poolWeightClient) {
	c.Lock()
	defer c.Unlock()

	for _, cli := range c.clients {
		if cli.endpoint == addr {
			return
		}
	}
	heap.Push(c, weightClient)

	if c.Len() == minHeapSize {
		heap.Init(c)
	}
}

// update clients with new addrs, remove the no use client
func (c *RpcClusterClient) clear(addrs []string) {
	c.Lock()
	var rm []*poolWeightClient
	for _, cli := range c.clients {
		var has_cli bool
		for _, addr := range addrs {
			if cli.endpoint == addr {
				has_cli = true
				break
			}
		}
		if !has_cli {
			rm = append(rm, cli)
		} else if cli.errcnt > 0 {
			/*
				if cli.weight >= errWeight*uint64(cli.errcnt) {
					cli.weight -= errWeight * uint64(cli.errcnt)
					cli.errcnt = 0
					if c.Len() >= minHeapSize {
						// cli will and only up, so it's ok here.
						heap.Fix(c, cli.index)
					}
				}
			*/
		}
	}

	for _, cli := range rm {
		// p will up, down, or not move, so append it to rm list.
		c.Debugf("remove cli: %s", cli.endpoint)

		heap.Remove(c, cli.index)
		cli.pool.Close()
	}
	c.Unlock()
}

func (c *RpcClusterClient) get() *poolWeightClient {
	c.Lock()
	defer c.Unlock()

	size := c.Len()
	if size == 0 {
		return nil
	}

	if size < minHeapSize {
		var index int = 0
		for i := 1; i < size; i++ {
			if c.Less(i, index) {
				index = i
			}
		}

		return c.clients[index]
	}

	client := heap.Pop(c).(*poolWeightClient)
	heap.Push(c, client)
	return client
}

func (c *RpcClusterClient) use(client *poolWeightClient) {
	c.Lock()
	client.weight++
	if c.Len() >= minHeapSize {
		heap.Fix(c, client.index)
	}
	client.stats.UseCount++
	c.Unlock()
}

func (c *RpcClusterClient) done(client *poolWeightClient) {
	/*
		c.Lock()
		if client.weight > 0 {
			client.weight--
		}
		if c.Len() >= minHeapSize {
			heap.Fix(c, client.index)
		}
		c.Unlock()
	*/
}

func (c *RpcClusterClient) occurErr(client *poolWeightClient, err error) {
	c.Lock()
	if nil != err {
		client.weight += errWeight
		client.errcnt++
		if c.Len() >= minHeapSize {
			heap.Fix(c, client.index)
		}

		client.stats.ErrorCount++
	} else {
		/*
			if client.errcnt > 0 {
				if client.weight >= errWeight {
					client.weight -= errWeight
				}
				client.errcnt--
				if c.Len() >= minHeapSize {
					heap.Fix(c, client.index)
				}
			}
		*/
	}
	c.Unlock()
}

func (c *RpcClusterClient) updateClientAddr() {
	addr := strings.Split(c.addr, ":")
	addrs, err := net.LookupHost(addr[0])
	if nil != err {
		c.Errorf("lookup host err: ", c.addr, err)
		log.Println("lookup host err: ", c.addr, err)
		return
	}
	// only ipv4
	var ips []string
	for _, s := range addrs {
		ip := net.ParseIP(s)
		if ip != nil && len(ip.To4()) == net.IPv4len {
			ips = append(ips, s)
		}
	}

	c.endpoints = ips

	//统计打印
	c.RLock()
	c.Debugf("############clients stats %s#############", c.addr)

	for i := range c.clients {
		c.Debugf("## ip :%s  use count :%d  index : %d  err total :%d  err peroid :%d  weights : %d",
			c.clients[i].endpoint, c.clients[i].stats.UseCount, c.clients[i].index, c.clients[i].stats.ErrorCount,
			c.clients[i].errcnt, c.clients[i].weight)
	}
	c.RUnlock()

	c.clear(ips)

	for i := range ips {
		if !c.exist(ips[i]) {
			newPoolWeightClient := &poolWeightClient{
				endpoint:      ips[i],
				port:          addr[1],
				clusterClient: c,
			}
			newPoolWeightClient.pool = NewPool(newPoolWeightClient.Connect, 100, 100, 1000*time.Second, true)
			c.add(ips[i], newPoolWeightClient)

			c.Debugf("add cli: %s", newPoolWeightClient.endpoint)
		}
	}

	c.RLock()
	if c.Len() == 0 {
		c.Errorf("cluster has no client to use")
		log.Println("cluster has no client to use")
	}
	c.RUnlock()
}

func (c *RpcClusterClient) CallTimeout(serviceMethod string, args interface{}, reply interface{}) error {
	var err error

	client := c.get()
	if nil == client {
		c.updateClientAddr()
		c.Errorf("call %s nil client", serviceMethod)
		return fmt.Errorf("nil client")
	}

	c.use(client)

	rpc_client, err := client.pool.Get()
	if err != nil {
		c.Errorf("get %s rpc client error :%v", c.addr, err)

		c.done(client)
		c.occurErr(client, err)
		return err
	}
	err = rpc_client.CallTimeout(serviceMethod, args, reply, c.retry, c.timeout)
	if err != nil {
		c.Errorf("call %s rpc client err :%s,%v", c.addr, serviceMethod, err)
	}
	rpc_client.Close()
	c.done(client)
	c.occurErr(client, err)

	return err
}
