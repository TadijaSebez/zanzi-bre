package server

import (
	"net/http"
	"placeholder/back/core"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

func login(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func register(c echo.Context) error {
	cc := c.(*CustomContext)
	s := cc.Server

	var dto core.UserDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	if s.Db.CheckUser(dto) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "user with that email already exists",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 5)
	if err != nil {
		return err
	}

	u := &core.User{}
	u.Password = string(hash)
	u.Email = dto.Email
	u.Name = dto.Name

	_, err = s.Db.InsertUser(*u)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.NoContent(200)
}
