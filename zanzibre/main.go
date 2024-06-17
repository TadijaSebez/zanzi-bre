package main

import (
	"flag"
	"log"
	"placeholder/zanzibar/server"
)

func main() {
	port := flag.Int("p", 6969, "Port number in range [1000, 65535]. Default is 6969")
	ip := flag.String("ip", "127.0.0.1", "IP address on which to serve. Default is 127.0.0.1")
	dbPath := flag.String("db", "zanzibase", "Path to the LevelDB folder. Default is ./zanzibase")

	flag.Parse()

	server, err := server.New(*ip, *port, *dbPath)

	if err != nil {
		log.Fatal(err)
	}

	server.Serve()
}
