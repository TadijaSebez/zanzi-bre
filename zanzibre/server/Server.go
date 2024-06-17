package server

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Port   string
	Router *echo.Echo
}

func newRouter() *echo.Echo {
	e := echo.New()

	e.GET("/put", dbPutTest)
	e.GET("/get", dbGetTest)

	return e
}

func New(port int) (*Server, error) {
	if port < 1000 || port > 65535 {
		return nil, fmt.Errorf("invalid port value")
	}

	router := newRouter()

	server := &Server{
		Port:   strconv.Itoa(port),
		Router: router,
	}

	return server, nil
}

func (s *Server) Serve() {
	s.Router.Start("localhost:" + s.Port)
}
