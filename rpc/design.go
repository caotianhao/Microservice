package main

import (
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//----------------------------------------------服务端封装

// DesignInterface 要求服务端在注册 rpc 对象时
// 能让编译器在编译时就检测对象是否合法
// 创建接口，在接口中定义方法原型
type DesignInterface interface {
	AddHelloAtHead(string, *string) error
}

// RegisterService 调用该方法时需要给 i 传参
// 参数应该是实现了 AddHelloAtHead 方法的类对象
// 需要联合编译，即 go run server.go design.go
func RegisterService(i DesignInterface) {
	err := rpc.RegisterName("myRPCSevName", i)
	if err != nil {
		log.Fatal("注册 rpc 服务失败！", err)
	}
}

//----------------------------------------------客户端封装

type DesignClient struct {
	c *rpc.Client
}

// InitClient 由于使用 c 调用了 Call，因此需要初始化
func InitClient(addr string) DesignClient {
	conn, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		log.Fatal("design jsonrpc.Dial err =", err)
	}
	return DesignClient{conn}
}

func (dc *DesignClient) AddHelloAtHead(a string, b *string) error {
	return dc.c.Call("myRPCSevName.AddHelloAtHead", a, &b)
}
