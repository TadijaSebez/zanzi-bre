package server

import (
	"fmt"
	"log"
	"placeholder/back/repository"
	"strconv"

	"github.com/labstack/echo/v4"
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

	// {content, title}
	e.POST("/note", save)

	// {noteId, userId, permission}
	e.POST("/share", share)

	// {noteId, userId}
	e.POST("/unshare", unshare)

	return e
}

func (s *Server) Serve() {
	router := newRouter(s)
	url := fmt.Sprintf("%s:%s", s.Ip, s.Port)
	log.Fatal(router.Start(url))
}
