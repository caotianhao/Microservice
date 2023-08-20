package main

import (
	"github.com/hashicorp/consul/api"
	"log"
)

func main() {
	// 1
	consulConfig := api.DefaultConfig()

	// 2
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal("api.NewClient() err =", err)
	}

	// 3
	err = consulClient.Agent().ServiceDeregister("cth01")
	if err != nil {
		log.Fatal("Deregister err =", err)
	}
}
