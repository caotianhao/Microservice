package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

type Hello struct {
	// empty
}

func (h *Hello) AddHelloAtHead(name string, r *string) error {
	*r = "Hello " + name + "!"
	return nil
}

func main() {
	// 1. 注册 rpc 服务
	//err := rpc.RegisterName("myRPCSevName", new(Hello))
	//if err != nil {
	//	log.Fatal("注册 rpc 服务失败！", err)
	//}
	RegisterService(new(Hello))

	// 2. 设置监听
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("net.Listen 失败！", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal("listener.Close() 失败！", err)
		}
	}(listener)
	fmt.Println("开始监听......")

	// 3. 接收数据
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("listener.Accept() 失败！", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("server.conn.Close() 失败！", err)
		}
	}(conn)

	// 4. rpc 的调用，绑定服务
	jsonrpc.ServeConn(conn)
}
