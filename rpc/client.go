package main

import (
	"fmt"
	"log"
)

func main() {
	// 1. rpc 连接服务器
	//conn, err := jsonrpc.Dial("tcp", ":8080")
	//if err != nil {
	//	log.Fatal("rpc.Dial err =", err)
	//}
	//defer func(client *rpc.Client) {
	//	err := client.Close()
	//	if err != nil {
	//		log.Fatal("client.Close() err =", err)
	//	}
	//}(conn)
	cli := InitClient(":8080")

	// 2. 远程函数调用
	//var result string
	//err = conn.Call("myRPCSevName.AddHelloAtHead", "rpc World", &result)
	//if err != nil {
	//	log.Fatal("client.Call() err =", err)
	//}
	//fmt.Println("成功！", result)
	var result string
	err := cli.AddHelloAtHead("design success", &result)
	if err != nil {
		log.Fatal("design cli.AddHelloAtHead err =", err)
	}
	fmt.Println(result)
}
