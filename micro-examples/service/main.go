/**
 * @author liangbo
 * @email  liangbo@codoon.com
 * @date   2018/5/29 下午6:25
 */
package main

import (
	"context"
	proto "gopher/micro-examples/service/proto"
	"github.com/micro/go-micro"
	"fmt"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, resp *proto.HelloResponse) error {
	resp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

