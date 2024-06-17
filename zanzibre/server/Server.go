package server

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/syndtr/goleveldb/leveldb"
)

type Server struct {
	Port string
}

type CustomContext struct {
	echo.Context
	Server *Server
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

	e.GET("/put", dbPutTest)
	e.GET("/get", dbGetTest)

	return e
}

func New(port int) (*Server, error) {
	if port < 1000 || port > 65535 {
		return nil, fmt.Errorf("invalid port value")
	}

	server := &Server{
		Port: strconv.Itoa(port),
	}

	return server, nil
}

func (s *Server) Serve() {
	router := newRouter(s)
	router.Start("localhost:" + s.Port)
}

func dbPut(key, value []byte) error {
	// TODO: Load db name from file?
	db, err := leveldb.OpenFile("zanzibase", nil)

	if err != nil {
		return err
	}

	defer db.Close()

	err = db.Put(key, value, nil)

	return err
}

func dbGet(key []byte) ([]byte, error) {
	// TODO: Load db name from file?
	db, err := leveldb.OpenFile("zanzibase", nil)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	value, err := db.Get(key, nil)

	return value, err
}
