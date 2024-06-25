package server

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"placeholder/zanzibar/core"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/syndtr/goleveldb/leveldb"
)

type Server struct {
	Port   string
	Ip     string
	DbPath string
	Engine *core.Engine
	PyPath string
}

type CustomContext struct {
	echo.Context
	Server *Server
}

type Acl struct {
	Object   string `json:"object"`
	Relation string `json:"relation"`
	User     string `json:"user"`
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
	e.GET("/acl/check", getAcl)
	e.POST("/acl/delete", delAcl)
	e.POST("/namespace", addNamespace)

	return e
}

func New(ip string, port int, dbPath string, engine *core.Engine, pyPath string) (*Server, error) {
	if port < 1000 || port > 65535 {
		return nil, fmt.Errorf("invalid port value")
	}

	server := &Server{
		Port:   strconv.Itoa(port),
		Ip:     ip,
		DbPath: dbPath,
		Engine: engine,
		PyPath: pyPath,
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

func (s *Server) dbDel(key []byte) error {
	db, err := leveldb.OpenFile(s.DbPath, nil)

	if err != nil {
		return err
	}

	defer db.Close()

	err = db.Delete(key, nil)

	return err
}

func (s *Server) checkAcl(object, relation, user string) bool {
	namespace := strings.Split(object, ":")[0]
	model := findNamespace(s.Engine.Template, namespace)

	if model == nil {
		return false
	}

	r := findRelation(model, relation)
	if r == nil {
		return false
	}

	return s.evaluateRelation(model, r, object, user)
}

func (s *Server) evaluateRelation(template *core.Model, relation *core.Relation, object, user string) bool {
	// If additional info is 0 that means the relation looks like { name: "owner" }
	if len(relation.AdditionalInfo) == 0 {
		return s.isInDb(object, relation.Name, user)
	}

	for _, info := range relation.AdditionalInfo {
		childBoolArr := make([]bool, 0)
		for _, child := range info.Children {
			if child.RelationName == "this" {
				// If this, add true only if in database
				childBoolArr = append(childBoolArr, s.isInDb(object, relation.Name, user))
			} else {
				// for example for { relation: "editor" }
				// it has to calculate its computed userset
				childRelation := findRelation(template, child.RelationName)
				if childRelation != nil {
					childBoolArr = append(childBoolArr, s.evaluateRelation(template, childRelation, object, user))
				} else {
					childBoolArr = append(childBoolArr, false)
				}
			}
		}

		// Evaluateing bools and returning false if false
		// because all additional infos (unions and intersections) must
		// return true in order for authorization to be approved
		if !evaluateBoolArray(childBoolArr, info.Type) {
			return false
		}
	}

	return true
}

func evaluateBoolArray(boolArr []bool, relationType string) bool {
	// TODO: These can be enums but Go doesn't really have a nice way for that,
	// will have to change when implementing "Exclusion"
	if relationType == "Union" {
		for _, b := range boolArr {
			if b {
				return true
			}
		}
		return false
	} else if relationType == "Intersection" {
		for _, b := range boolArr {
			if !b {
				return false
			}
		}
		return true
	}
	return false
}

func findRelation(template *core.Model, name string) *core.Relation {
	for _, r := range template.Relations {
		if r.Name == name {
			return r
		}
	}
	return nil
}

func findNamespace(template []*core.Model, namespace string) *core.Model {
	for _, m := range template {
		if m.Namespace == namespace {
			return m
		}
	}
	return nil
}

func (s *Server) isInDb(object, relation, user string) bool {
	query := fmt.Sprintf("%s#%s@%s", object, relation, user)
	_, err := s.dbGet([]byte(query))
	return err == nil
}

func (s *Server) addNamespace() (string, error) {
	//grammar := "--grammar ./parser/grammar.tx"
	//model := "--model temp.ent"
	cmd := exec.Command("python", s.PyPath, "--grammar", "./parser/grammar.tx", "--model", "temp.ent")

	var out bytes.Buffer
	cmd.Stdout = &out

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	fmt.Println(out.String())
	fmt.Println(stderr.String())
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
