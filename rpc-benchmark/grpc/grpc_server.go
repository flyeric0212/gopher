/**
 * @author liangbo
 * @email  liangbo@codoon.com
 * @date   2018/2/7 下午4:28
 */
package main

import (
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"runtime"
	"time"
)

type Hello struct{}

func (t *Hello) Say(ctx context.Context, args *BenchmarkMessage) (reply *BenchmarkMessage, err error) {
	s := "OK"
	var i int32 = 100
	args.Field1 = s
	args.Field2 = i
	if *delay > 0 {
		time.Sleep(*delay)
	} else {
		runtime.Gosched()
	}
	return args, nil
}

var (
	host  = flag.String("s", "127.0.0.1:8973", "listened ip and port")
	delay = flag.Duration("delay", 0, "delay to mock business processing")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	RegisterHelloServer(s, &Hello{})
	s.Serve(lis)
}
