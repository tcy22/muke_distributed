package main

import (
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
	rpcDemo "muke_distributed/crawier/rpc"
)

func main() {
	//1、服务端需要注册可以被远程访问的对象，之后这个对象就可以被远程访问了
	rpc.Register(rpcDemo.DemoService{})
	//2、创建网络监听器,等待发过来的请求
	listener,err := net.Listen("tcp",":2222")
	if err != nil {
		panic(err)
	}

	for {
		//3、获取发送来的请求
		//Accept接收监听器l获取的连接
		conn,err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v",err)
			continue
		}
		//4、异步处理发送来的请求
		go jsonrpc.ServeConn(conn)
	}
}
