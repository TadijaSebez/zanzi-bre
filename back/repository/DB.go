package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	Host       string
	Port       int
	User       string
	Password   string
	DbName     string
	ConnString string
}

func New() *DB {
	db := &DB{
		Host:     "localhost",
		Port:     5432,
		User:     "user",
		Password: "user",
		DbName:   "back",
	}
	db.ConnString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.DbName)
	return db
}

func (d *DB) CreateTables() {
	db, err := sql.Open("postgres", d.ConnString)

	if err != nil {
		panic(err)
	}

	defer db.Close()
	table := "Note(id SERIAL PRIMARY KEY, title VARCHAR(255) NOT NULL, content TEXT NOT NULL);"
	_, err = db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s", table))

	if err != nil {
		log.Fatal("failed to create table with error message ", err.Error())
	}
}
