package main

import (
	"Microservice/grpc_exercise2/pe"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// 1. 连接 grpc 服务
	// 后面的参数是为了保证安全性，不加有时候会报错
	grpcConn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("grpc.Dial err =", err)
	}
	defer func(grpcConn *grpc.ClientConn) {
		err := grpcConn.Close()
		if err != nil {
			log.Fatal("grpcConn.Close() err =", err)
		}
	}(grpcConn)

	// 2. 初始化 grpc 客户端
	grpcClient := pe.NewInfoClient(grpcConn)

	// 创建并初始化 City 对象
	var MyCity pe.City
	MyCity.Name = "Chongqing"
	MyCity.ID = 1

	// 3. 调用远程服务
	// 正常有上下文参数传参时直接传入，没有就传入空对象
	r, err := grpcClient.Show(context.TODO(), &MyCity)
	if err != nil {
		log.Fatal("grpcClient.Show() err =", err)
	}
	fmt.Println(r)
}
