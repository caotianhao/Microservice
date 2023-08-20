package main

import (
	"Microservice/grpc_exercise4/g4"
	"context"
	"google.golang.org/grpc"
	"net"
)

type BBB struct {
}

func (b *BBB) Zeno(_ context.Context, m *g4.Mota) (*g4.Mota, error) {
	m.Name += "斯莉英雄传"
	m.ID = 99999
	return m, nil
}

func main() {
	// 1
	gs := grpc.NewServer()

	// 2
	g4.RegisterDestroyServer(gs, new(BBB))

	// 3
	listener, _ := net.Listen("tcp", ":8080")
	defer func(listener net.Listener) {
		_ = listener.Close()
	}(listener)

	// 4
	_ = gs.Serve(listener)
}
