package server

import (
	"fmt"
	"log"
	"placeholder/back/core"
	"placeholder/back/repository"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Port string
	Ip   string
	Db   *repository.DB
}

type CustomContext struct {
	echo.Context
	Server *Server
}

func New(port int, ip string) (*Server, error) {
	if port < 1000 || port > 65535 {
		return nil, fmt.Errorf("invalid port value")
	}

	db := repository.New()
	err := db.CreateTables()

	if err != nil {
		return nil, fmt.Errorf("couldn't connect to the database")
	}

	s := &Server{
		Port: strconv.Itoa(port),
		Ip:   ip,
		Db:   db,
	}

	return s, nil
}

func newRouter(s *Server) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	// Middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{
				Context: c,
				Server:  s,
			}
			return next(cc)
		}
	})

	// []Note
	e.GET("/note:userId", getAll)

	// Note
	e.POST("/note", save)

	// {noteId, userId, permission}
	e.POST("/share", share)

	// {noteId, userId}
	e.POST("/unshare", unshare)

	// {email, password}
	e.POST("/login", login)

	// {name, email, password}
	e.POST("/register", register)

	return e
}

func (s *Server) Serve() {
	router := newRouter(s)
	url := fmt.Sprintf("%s:%s", s.Ip, s.Port)
	log.Fatal(router.Start(url))
}

func (s *Server) Save(dto core.Note) (*core.Note, error) {
	if dto.Id == -1 {
		return s.Db.Insert(dto)
	}

	return s.Db.Update(dto)
}
