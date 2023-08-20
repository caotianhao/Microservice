package main

import (
	"Microservice/grpc_exercise2/pe"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type CityStruct struct {
	// empty
}

func (cs *CityStruct) Show(_ context.Context, c *pe.City) (*pe.City, error) {
	c.Name += "Shi"
	c.ID += 10000
	return c, nil
}

func main() {
	// 1. 初始 grpc 对象
	gs := grpc.NewServer()

	// 2. 注册服务
	pe.RegisterInfoServer(gs, new(CityStruct))

	// 3. 设置监听
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("net.Listen err =", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal("listener.Close() err =", err)
		}
	}(listener)

	// 4. 启动服务
	err = gs.Serve(listener)
	if err != nil {
		log.Fatal("gs.Serve err =", err)
	}
}
