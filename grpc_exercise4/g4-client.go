package main

import (
	"Microservice/grpc_exercise4/g4"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1
	conn, _ := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	// 2
	var QiKeng g4.Mota
	QiKeng.Name = "弃坑立粽传"
	QiKeng.ID = 9
	cli := g4.NewDestroyClient(conn)

	// 3
	zeno, _ := cli.Zeno(context.TODO(), &QiKeng)
	fmt.Println(zeno)
}
