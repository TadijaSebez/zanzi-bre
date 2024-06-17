package main

import (
	"log"
	"placeholder/zanzibar/server"
)

func main() {
	server, err := server.New(6969)

	if err != nil {
		log.Fatal(err)
	}

	server.Serve()
}
