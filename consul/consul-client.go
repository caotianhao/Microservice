package main

import (
	"Microservice/consul/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
)

func main() {
	// 1. 初始化 consul 配置
	consulConfig := api.DefaultConfig()

	// 2. 创建 consul 对象（可以重新指定 consul 属性，例如 ip 和端口，也可以使用默认）
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal("api.NewClient() err =", err)
	}

	// 3. 服务发现：从 consul 上获取健康的服务
	// func (h *Health) Service(service, tag string, passingOnly bool, q *QueryOptions)
	// 参数：
	// 		service: 服务名。 -- 注册服务时，指定该string，即Name字段
	// 		tag：外名/别名。 如果有多个， 任选一个
	// 		passingOnly：是否通过健康检查。 true
	//	 	q：查询参数。 通常传 nil
	// 返回值： ([]*ServiceEntry, *QueryMeta, error)
	// 		ServiceEntry： 存储服务的切片。
	// 		QueryMeta：额外查询返回值。 nil
	// 		error： 错误信息
	services, _, err := consulClient.Health().Service("HelloService", "testHello", true, nil)
	if err != nil {
		log.Fatal("consulClient.Health().Service err =", err)
	}
	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)

	// 以下为 grpc 服务远程调用 ======================================================
	//conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// 换为使用 consul 获取到的 addr
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("grpc.Dial() err =", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("conn.Close() err =", err)
		}
	}(conn)

	var p1 pb.Person
	p1.Age = 9999
	p1.Name = "CaoTh"
	cli := pb.NewHelloClient(conn)

	r, err := cli.SayHello(context.TODO(), &p1)
	if err != nil {
		log.Fatal("cli.SayHello() err =", err)
	}
	fmt.Println(r)
}
