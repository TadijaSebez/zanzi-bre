package main

import (
	"flag"
	"log"
	"placeholder/back/server"
)

func main() {
	port := flag.Int("p", 1337, "Port number in range [1000, 65535].")
	ip := flag.String("ip", "127.0.0.1", "IP address on which to serve.")
	zanzibarPort := flag.Int("Zp", 1337, "Port number of the zanzi-bre service in range [1000, 65535].")
	zanzibarIp := flag.String("Zip", "127.0.0.1", "IP address of the zanzi-bre service.")

	flag.Parse()

	s, err := server.New(*port, *ip, *zanzibarPort, *zanzibarIp)

	if err != nil {
		log.Fatal(err)
	}

	s.Serve()
}
