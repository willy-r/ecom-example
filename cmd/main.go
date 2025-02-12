package main

import (
	"log"

	"github.com/willy-r/ecom-example/cmd/api"
)

func main() {
	server := api.NewApiServer(":8080", nil)

	if err := server.Start(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
