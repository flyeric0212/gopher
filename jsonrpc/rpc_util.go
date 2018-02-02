package jsonrpc

import (
	"backend/jsonrpc"
	"backend/rpc"
	"net"
	"time"
)

const (
	RetryTime = 2
)

type Log interface {
	Error(format string, args ...interface{})
	Info(format string, args ...interface{})
	Notice(format string, args ...interface{})
	Debug(format string, args ...interface{})
}

type RpcClient struct {
	rpc_client *rpc.Client
	Addr       string
	Net        string
	name       string
	func_map   map[string]string
	logger     rpc.Log
	//pool        *rpc.Pool
	poolCluster *rpc.RpcClusterClient
}

func NewRpcClient(addr, net string, func_map map[string]string, name string,
	logger rpc.Log) (*RpcClient, error) {
	var err error
	client := &RpcClient{}
	client.Addr = addr
	client.Net = net
	client.func_map = func_map
	client.name = name
	client.logger = logger
	//client.pool = rpc.NewPool(client.Connect, 100, 100, 1000*time.Second, true)
	client.poolCluster = rpc.NewRpcClusterClient(jsonrpc.NewClient, addr, logger, 2*time.Second, RetryTime)

	return client, err
}

func (client *RpcClient) Connect() (*rpc.Client, error) {
	conn, err := net.DialTimeout(client.Net, client.Addr, 2*time.Second)
	if err != nil {
		if nil != client.logger {
			client.logger.Error("get %s rpc client error :%v", client.name, err)
		}
		return nil, err
	}

	rpc_client := jsonrpc.NewClient(conn)

	return rpc_client, nil
}

func (client *RpcClient) Call(method string, args interface{}, reply interface{}) error {
	var err error
	if nil != client.logger {
		client.logger.Debug("call rpc : %s, %v", method, args)
	}

	err = client.poolCluster.CallTimeout(client.func_map[method], args, reply)
	if err != nil {
		if nil != client.logger {
			client.logger.Error("call %s rpc client err :%s,%v", client.name, method, err)
		}
	}

	return err
}

func (client *RpcClient) CallMethod(method string, args interface{}, reply interface{}) error {
	var err error
	if err != nil {
		if nil != client.logger {
			client.logger.Error("get %s rpc client error :%v", client.name, err)
		}
		return err
	}

	err = client.poolCluster.CallTimeout(method, args, reply)
	if err != nil {
		if nil != client.logger {
			client.logger.Error("call %s rpc client err :%s,%v", client.name, method, err)
		}
	}

	return err
}

func (client *RpcClient) DirectCall(method string, args interface{}, reply interface{}) error {

	err := client.poolCluster.CallTimeout(client.func_map[method], args, reply)
	if err != nil {
		client.logger.Error("call %s rpc client err :%s,%v", client.name, method, err)
	}

	return err
}
