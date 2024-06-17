package server

import (
	"fmt"
	"net/http"

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

	if o == "" || r == "" || u == "" {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "expected query parameters object, relation, user.",
		})
	}

	key := fmt.Sprintf("%s#%s@%s", o, r, u)
	_, err := s.dbGet([]byte(key))

	var value string = "true"
	if err != nil {
		value = "false"
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"authorized": value,
	})
}
