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
	// Note
	noteTable := "Note(id SERIAL PRIMARY KEY, title VARCHAR(255) NOT NULL, content TEXT NOT NULL);"
	_, err = db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s", noteTable))

	if err != nil {
		return err
	}

	// User
	userTable := "Users(id SERIAL PRIMARY KEY, email VARCHAR(255) NOT NULL, password VARCHAR(255) NOT NULL, name VARCHAR(255) NOT NULL);"
	_, err = db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s", userTable))

	if err != nil {
		return err
	}

	// User_Note
	userNoteTable := `
    Users_Note (
        user_id INT NOT NULL,
        note_id INT NOT NULL,
        PRIMARY KEY (user_id, note_id),
        FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
        FOREIGN KEY (note_id) REFERENCES Note(id) ON DELETE CASCADE
    );
`
	_, err = db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s", userNoteTable))

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

	note.NoteId = id
	return &note, err
}

func (d *DB) InsertUser(user core.User) (*core.User, error) {
	db, err := sql.Open("postgres", d.ConnString)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Users (email, password, name) VALUES ($1, $2, $3) RETURNING ID")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var id int
	err = stmt.QueryRow(user.Email, user.Password, user.Name).Scan(&id)

	user.Id = id
	return &user, err
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

	_, err = stmt.Exec(note.Content, note.Title, note.NoteId)

	return &note, err
}

func (d *DB) CheckUser(user core.UserDTO) bool {
	db, err := sql.Open("postgres", d.ConnString)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := `SELECT email, password, name FROM Users WHERE email = $1`
	rows, err := db.Query(query, user.Email)

	if err != nil {
		return false
	}

	defer rows.Close()

	return rows.Next()
}

func (d *DB) LoginUser(dto core.LoginDTO) (*core.User, error) {
	db, err := sql.Open("postgres", d.ConnString)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := `SELECT id, email, password, name FROM Users WHERE email = $1`
	rows, err := db.Query(query, dto.Email)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("user not found")
	}

	var id int
	var email, password, name string
	err = rows.Scan(&id, &email, &password, &name)

	u := &core.User{
		Id:       id,
		Name:     name,
		Password: password,
		Email:    email,
	}

	return u, err
}

func (d *DB) GetAllNotes() ([]*core.Note, error) {
	db, err := sql.Open("postgres", d.ConnString)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := `SELECT id, content, title FROM Note`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var notes []*core.Note
	for rows.Next() {
		var note core.Note
		if err := rows.Scan(&note.NoteId, &note.Content, &note.Title); err != nil {
			return nil, err
		}
		notes = append(notes, &note)
	}

	return notes, nil
}

func (d *DB) FindUser(email string) *core.SharedUserDTO {
	db, err := sql.Open("postgres", d.ConnString)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := `SELECT id, email, name FROM Users WHERE email = $1`
	rows, err := db.Query(query, email)

	if err != nil {
		return nil
	}

	defer rows.Close()

	if rows.Next() {
		u := core.SharedUserDTO{}
		if err := rows.Scan(&u.Id, &u.Email, &u.Name); err != nil {
			return nil
		}
		return &u
	}

	return nil
}

func (d *DB) GetAllUsers() ([]*core.User, error) {
	db, err := sql.Open("postgres", d.ConnString)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := `SELECT id, email, name FROM Users`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*core.User
	for rows.Next() {
		var user core.User
		if err := rows.Scan(&user.Id, &user.Email, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
