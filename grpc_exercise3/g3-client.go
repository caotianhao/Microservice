package main

import (
	"Microservice/grpc_exercise3/g3test"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	gc := g3test.NewGameInfoClient(conn)

	var myGame g3test.Game
	myGame.ID = 99
	myGame.Name = "Mota"

	game, err := gc.Test3(context.TODO(), &myGame)
	if err != nil {
		return
	}
	fmt.Println(game)
}
