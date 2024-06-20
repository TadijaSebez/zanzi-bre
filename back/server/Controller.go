package server

import (
	"net/http"
	"placeholder/back/core"

	"github.com/labstack/echo/v4"
)

func getAll(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func save(c echo.Context) error {
	cc := c.(*CustomContext)
	s := cc.Server

	var dto core.Note
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	note, err := s.Save(dto)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, note)
}

func share(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func unshare(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
