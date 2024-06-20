package repository

import (
	"database/sql"
	"fmt"
	"placeholder/back/core"

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

func (d *DB) CreateTables() error {
	db, err := sql.Open("postgres", d.ConnString)

	if err != nil {
		return err
	}

	defer db.Close()
	table := "Note(id SERIAL PRIMARY KEY, title VARCHAR(255) NOT NULL, content TEXT NOT NULL);"
	_, err = db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s", table))

	return err
}

func (d *DB) Insert(note core.Note) (*core.Note, error) {
	db, err := sql.Open("postgres", d.ConnString)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Note (content, title) VALUES ($1, $2) RETURNING ID")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var id int
	err = stmt.QueryRow(note.Content, note.Title).Scan(&id)

	note.Id = id
	return &note, err
}

func (d *DB) Update(note core.Note) (*core.Note, error) {
	db, err := sql.Open("postgres", d.ConnString)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("UPDATE Note SET content=$1, title=$2 WHERE id=$3")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(note.Content, note.Title, note.Id)

	return &note, err
}
