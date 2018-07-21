package main

import (
	"net"
	"net/rpc/jsonrpc"
	"fmt"
	"muke_distributed/crawier/rpc"
)

//客户端要呼叫服务端
func main()  {
	//1、创建连接
	//Dial在指定的网络和地址与RPC服务端连接
	conn,err := net.Dial("tcp",":2222")
	if err != nil {
		panic(err)
	}
	//2、创建客户端
	client:=jsonrpc.NewClient(conn)

	var result float64
	//3、客户端通过连接向服务端发送请求
	err = client.Call("DemoService.Div",
		rpc.Args{10,3},&result)
	fmt.Println(result,err)
}
