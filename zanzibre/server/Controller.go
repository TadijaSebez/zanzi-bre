package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func putAcl(c echo.Context) error {
	cc := c.(*CustomContext)
	s := cc.Server

	var acl Acl
	if err := c.Bind(&acl); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	key := fmt.Sprintf("%s#%s@%s", acl.Object, acl.Relation, acl.User)
	fmt.Printf("[*] Putting: %s\n", key)

	err := s.dbPut([]byte(key), []byte(""))

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.String(http.StatusOK, key)
}

func getAcl(c echo.Context) error {
	cc := c.(*CustomContext)
	s := cc.Server

	o := c.QueryParam("object")
	r := c.QueryParam("relation")
	u := c.QueryParam("user")

	fmt.Printf("[*] Checking: %s#%s@%s\n", o, r, u)

	if o == "" || r == "" || u == "" {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "expected query parameters object, relation, user.",
		})
	}

	isOk := s.checkAcl(o, r, u)
	value := strconv.FormatBool(isOk)

	return c.JSON(http.StatusOK, map[string]string{
		"authorized": string(value),
	})
}

func delAcl(c echo.Context) error {
	cc := c.(*CustomContext)
	s := cc.Server

	var acl Acl
	if err := c.Bind(&acl); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	key := fmt.Sprintf("%s#%s@%s", acl.Object, acl.Relation, acl.User)
	fmt.Printf("[*] Deleting: %s\n", key)

	err := s.dbDel([]byte(key))

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.String(http.StatusOK, key)
}

func addNamespace(c echo.Context) error {
	cc := c.(*CustomContext)
	s := cc.Server

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	defer src.Close()

	dst, err := os.Create("temp.ent")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	template, err := s.addNamespace()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	s.Engine.AddTemplate(template)

	return c.NoContent(200)
}
