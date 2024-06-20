package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getAll(c echo.Context) error {
	//cc := c.(CustomContext)
	//s := cc.Server

	return c.String(http.StatusOK, "Hello, World!")
}

func save(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func share(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func unshare(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
