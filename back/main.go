package main

import (
	"flag"
	"log"
	"placeholder/back/server"
)

func main() {
	port := flag.Int("p", 1337, "Port number in range [1000, 65535].")
	ip := flag.String("ip", "127.0.0.1", "IP address on which to serve.")

	flag.Parse()

	s, err := server.New(*port, *ip)

	if err != nil {
		log.Fatal(err)
	}

	s.Serve()
}
