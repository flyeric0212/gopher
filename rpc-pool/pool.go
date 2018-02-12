/**
 * @author liangbo
 * @email  liangbo@codoon.com
 * @date   2018/2/12 下午6:17
 */
package pool

import (
	"container/list"
	"errors"
	"net/rpc"
	"sync"
	"time"
)

var (
	errPoolClosed = errors.New("rpc: connection pool closed")
	errConnClosed = errors.New("rpc: connection closed")
)

type Pool struct {
	MaxIdle     int
	MaxActive   int
	Wait        bool
	IdleTimeout time.Duration
	Dial        func() (*rpc.Client, error)
	mu          sync.Mutex
	cond        *sync.Cond
	closed      bool
	active      int
	idle        list.List
}
