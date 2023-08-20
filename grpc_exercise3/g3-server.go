package main

import (
	"Microservice/grpc_exercise3/g3test"
	"context"
	"google.golang.org/grpc"
	"net"
)

type AAA struct {
}

func (a *AAA) Test3(_ context.Context, g *g3test.Game) (*g3test.Game, error) {
	g.ID += 888
	g.Name += "Game."
	return g, nil
}

func main() {
	gs := grpc.NewServer()

	g3test.RegisterGameInfoServer(gs, new(AAA))

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			return
		}
	}(listener)

	err = gs.Serve(listener)
	if err != nil {
		return
	}
}
