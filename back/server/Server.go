package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"placeholder/back/core"
	"placeholder/back/repository"
	"strconv"
	"strings"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Port     string
	Ip       string
	Db       *repository.DB
	Zanzibar *core.Zanzibar
}

type CustomContext struct {
	echo.Context
	Server *Server
}

func New(port int, ip string, zanzibarPort int, zanzibarIp string) (*Server, error) {
	if port < 1000 || port > 65535 {
		return nil, fmt.Errorf("invalid port value")
	}

	if zanzibarPort < 1000 || zanzibarPort > 65535 {
		return nil, fmt.Errorf("invalid zanzibar port value")
	}

	db := repository.New()
	err := db.CreateTables()

	if err != nil {
		return nil, fmt.Errorf("couldn't connect to the database")
	}

	zanzibar := &core.Zanzibar{
		Ip:   zanzibarIp,
		Port: fmt.Sprintf(":%d", zanzibarPort),
	}

	s := &Server{
		Port:     strconv.Itoa(port),
		Ip:       ip,
		Db:       db,
		Zanzibar: zanzibar,
	}

	fmt.Println(s.Zanzibar.Ip, s.Zanzibar.Port)

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

	// {email, password}
	e.POST("/login", login)

	// {name, email, password}
	e.POST("/register", register)

	r := e.Group("")

	r.Use(echojwt.JWT([]byte("secretkey")))

	// []Note
	r.GET("/note", getAll)

	// Note
	r.POST("/note", save)

	// {noteId, userId, permission}
	r.POST("/share", share)

	// {noteId, userId}
	r.POST("/unshare", unshare)

	// {name, email, permission}
	r.GET("/share/:noteId", getShared)

	return e
}

func (s *Server) Serve() {
	router := newRouter(s)
	url := fmt.Sprintf("%s:%s", s.Ip, s.Port)
	log.Fatal(router.Start(url))
}

func (s *Server) Save(dto core.Note) (*core.Note, error) {
	if dto.NoteId == -1 {
		return s.Db.Insert(dto)
	}

	return s.Db.Update(dto)
}

func (s *Server) SendAcl(body []byte) error {
	url := fmt.Sprintf("http://%s%s%s", s.Zanzibar.Ip, s.Zanzibar.Port, core.PutEndpoint)
	_, err := s.SendRequest(http.MethodPost, url, body)
	return err
}

func (s *Server) DeleteAcl(body []byte) error {
	url := fmt.Sprintf("http://%s%s%s", s.Zanzibar.Ip, s.Zanzibar.Port, core.DelEndpoint)
	_, err := s.SendRequest(http.MethodPost, url, body)
	return err
}

func (s *Server) CheckAcl(object, relation, user string) (bool, error) {
	endpoint := fmt.Sprintf("http://%s%s%s", s.Zanzibar.Ip, s.Zanzibar.Port, core.CheckEndpoint)
	url := fmt.Sprintf("%s?object=note:%s&relation=%s&user=user:%s", endpoint, object, relation, user)
	resp, err := s.SendRequest(http.MethodGet, url, nil)
	b, _ := io.ReadAll(resp.Body)
	body := string(b[:])
	if strings.Contains(body, "true") {
		return true, err
	}
	return false, err
}

func (s *Server) SendRequest(method, url string, body []byte) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()

	//var reqBody *bytes.Buffer = nil
	//if body != nil {
	//	reqBody = bytes.NewBuffer(body)
	//}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	return response, err
}
