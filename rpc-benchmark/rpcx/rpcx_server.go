/**
 * @author liangbo
 * @email  liangbo@codoon.com
 * @date   2018/2/7 下午5:25
 */
package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"

	rlog "github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/server"
)

type Hello int

func (t *Hello) Say(ctx context.Context, args *BenchmarkMessage, reply *BenchmarkMessage) error {
	args.Field1 = "OK"
	args.Field2 = 100
	*reply = *args
	if *delay > 0 {
		time.Sleep(*delay)
	} else {
		runtime.Gosched()
	}
	return nil
}

var (
	host      = flag.String("s", "127.0.0.1:8974", "listened ip and port")
	delay     = flag.Duration("delay", 0, "delay to mock business processing")
	debugAddr = flag.String("d", "127.0.0.1:9984", "server ip and port")
)

func main() {
	flag.Parse()

	server.UsePool = true

	rlog.SetDummyLogger()

	go func() {
		log.Println(http.ListenAndServe(*debugAddr, nil))
	}()

	server := server.NewServer()
	server.RegisterName("Hello", new(Hello), "")
	server.Serve("tcp", *host)
}
