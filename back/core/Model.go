package core

type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ShareDTO struct {
	NoteId     int    `json:"noteId"`
	UserId     string `json:"userId"`
	Permission string `json:"permission"`
}

type Zanzibar struct {
	Ip   string
	Port string
}

const PutEndpoint string = "/acl"
const DelEndpoint string = "/acl/delete"
const CheckEndpoint string = "/acl/check"
