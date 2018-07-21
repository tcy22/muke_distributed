package rpcSupport

import (
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string,service interface{}) error{
	//1、服务端需要注册可以被远程访问的对象，之后这个对象就可以被远程访问了
	rpc.Register(service)
	//2、创建网络监听器,等待发过来的请求
	listener,err := net.Listen("tcp",host)
	if err != nil {
		return err
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
	return nil
}

func NewClient(host string) (*rpc.Client,error){
	conn,err := net.Dial("tcp",host)
	if err != nil {
		panic(err)
	}

	return jsonrpc.NewClient(conn),nil
}
