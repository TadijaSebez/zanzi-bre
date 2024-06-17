package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func dbPutTest(c echo.Context) error {
	key := "key"
	value := "value"

	err := dbPut([]byte(key), []byte(value))

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.String(http.StatusOK, fmt.Sprintf("%s: %s", key, value))
}

func dbGetTest(c echo.Context) error {
	key := "key"

	value, err := dbGet([]byte(key))

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.String(http.StatusOK, string(value))
}
