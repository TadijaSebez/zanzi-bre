package main

import (
	"flag"
	"log"
	"placeholder/zanzibar/core"
	"placeholder/zanzibar/server"
)

func main() {
	port := flag.Int("p", 6969, "Port number in range [1000, 65535].")
	ip := flag.String("ip", "127.0.0.1", "IP address on which to serve.")
	dbPath := flag.String("db", "zanzibase", "Path to the LevelDB folder.")
	templatePath := flag.String("t", "parser/template.json", "Path to the dsl template.")
	pyPath := flag.String("py", "./parser/parser.py", "Path to the python script.")

	flag.Parse()

	engine, err := core.NewEngine(*templatePath)

	if err != nil {
		log.Fatal(err)
	}

	server, err := server.New(*ip, *port, *dbPath, engine, *pyPath)

	if err != nil {
		log.Fatal(err)
	}

	engine.PrintTemplate()
	server.Serve()
}
