package main

import (
	"Microservice/consul/pb"
	"context"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Teacher struct {
}

func (t *Teacher) SayHello(_ context.Context, p *pb.Person) (*pb.Person, error) {
	p.Name = "Dr." + p.Name
	return p, nil
}

func main() {
	// consul.exe agent -server -bootstrap-expect 1 -data-dir D:\program\consul_1.16.0_windows_amd64\consul\ -node=n1 -bind=192.168.154.1 -ui -rejoin -config-dir=D:\program\consul_1.16.0_windows_amd64\consul.d\ -client 0.0.0.0
	// 真正用的时候 consul agent -dev 即可

	// 把 grpc 服务注册到 consul 上
	// 1. 初始化 consul 配置
	consulConfig := api.DefaultConfig()

	// 2. 创建 consul 对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal("api.NewClient() err =", err)
	}

	// 3. 告诉 consul 即将注册的服务的配置信息
	reg := api.AgentServiceRegistration{
		ID:      "cth01", // 注销时使用
		Tags:    []string{"testHello"},
		Name:    "HelloService", // 注册时使用
		Address: "127.0.0.1",
		Port:    8080,
		Check: &api.AgentServiceCheck{
			CheckID:  "consul check",
			TCP:      "127.0.0.1:8080",
			Timeout:  "1s",
			Interval: "5s",
		},
	}

	// 4. 注册 grpc 服务到 consul 上
	err = consulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal("consulClient.Agent().ServiceRegister(&reg) err =", err)
	}
	// 使用前需要 consul agent -dev 启动默认服务
	// 然后运行 server
	// 浏览器 127.0.0.1:8500

	// 以下为 grpc 服务远程调用 ======================================================
	gs := grpc.NewServer()

	pb.RegisterHelloServer(gs, new(Teacher))

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("net.Listen() err =", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal("listener.Close() err =", err)
		}
	}(listener)

	err = gs.Serve(listener)
	if err != nil {
		log.Fatal("gs.Serve() err =", err)
	}
}
