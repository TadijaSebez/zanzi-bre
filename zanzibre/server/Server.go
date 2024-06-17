package server

import (
	"fmt"
	"log"
	"placeholder/zanzibar/core"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/syndtr/goleveldb/leveldb"
)

type Server struct {
	Port   string
	Ip     string
	DbPath string
	Engine *core.Engine
}

type CustomContext struct {
	echo.Context
	Server *Server
}

type Acl struct {
	Object   string
	Relation string
	User     string
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

	e.POST("/acl", putAcl)
	//e.POST("/acl/check", getAcl)

	return e
}

func New(ip string, port int, dbPath string, engine *core.Engine) (*Server, error) {
	if port < 1000 || port > 65535 {
		return nil, fmt.Errorf("invalid port value")
	}

	server := &Server{
		Port:   strconv.Itoa(port),
		Ip:     ip,
		DbPath: dbPath,
		Engine: engine,
	}

	return server, nil
}

func (s *Server) Serve() {
	router := newRouter(s)
	url := fmt.Sprintf("%s:%s", s.Ip, s.Port)
	log.Fatal(router.Start(url))
}

func (s *Server) dbPut(key, value []byte) error {
	db, err := leveldb.OpenFile(s.DbPath, nil)

	if err != nil {
		return err
	}

	defer db.Close()

	err = db.Put(key, value, nil)

	return err
}

func (s *Server) dbGet(key []byte) ([]byte, error) {
	db, err := leveldb.OpenFile(s.DbPath, nil)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	value, err := db.Get(key, nil)

	return value, err
}
