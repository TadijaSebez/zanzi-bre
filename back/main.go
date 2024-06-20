package main

import "placeholder/back/repository"

func main() {
	db := repository.New()
	db.CreateTables()
}
