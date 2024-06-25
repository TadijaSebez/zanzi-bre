package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"placeholder/back/core"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func getAll(c echo.Context) error {
	cc := c.(*CustomContext)
	s := cc.Server

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	userId := strconv.FormatFloat(id, 'f', -1, 64)

	notes, err := s.Db.GetAllNotes()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	var ret []*core.UsersNoteDTO
	permissions := []string{"owner", "editor", "viewer"}
	for _, note := range notes {
		var userNote core.UsersNoteDTO
		for _, p := range permissions {
			if b, err := s.CheckAcl(strconv.Itoa(note.NoteId), p, userId); err == nil && b {
				userNote = core.UsersNoteDTO{
					NoteId:     note.NoteId,
					Content:    note.Content,
					Title:      note.Title,
					Permission: p,
				}
				ret = append(ret, &userNote)
				break
			}
		}
	}

	if ret == nil {
		return c.JSON(http.StatusOK, make([]string, 0))
	}

	return c.JSON(http.StatusOK, ret)
}

func save(c echo.Context) error {
	cc := c.(*CustomContext)
	s := cc.Server

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	userId := strconv.FormatFloat(id, 'f', -1, 64)

	var dto core.Note
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	var relation string
	if dto.NoteId != -1 {
		relation := "owner"
		if b, _ := s.CheckAcl(strconv.Itoa(dto.NoteId), relation, userId); !b {
			relation = "editor"
			if b, _ := s.CheckAcl(strconv.Itoa(dto.NoteId), relation, userId); !b {
				return fmt.Errorf("you are not the owner of this note")
			}
		}
	} else {
		relation = "owner"
	}

	note, err := s.Save(dto)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	payload := map[string]string{
		"object":   fmt.Sprintf("note:%d", note.NoteId),
		"relation": relation,
		"user":     fmt.Sprintf("user:%s", userId),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "error creating payload",
		})
	}

	if relation == "owner" {
		if err := s.SendAcl(jsonData); err != nil {
			return fmt.Errorf("failed to save the permission to zanzibar")
		}
	}

	return c.JSON(http.StatusOK, note)
}

func isOwner(c echo.Context, dto core.ShareDTO) error {
	cc := c.(*CustomContext)
	s := cc.Server

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id := claims["id"].(float64)
	idStr := strconv.FormatFloat(id, 'f', -1, 64)

	if b, err := s.CheckAcl(strconv.Itoa(dto.NoteId), "owner", idStr); err != nil && !b {
		return fmt.Errorf("you are not the owner of this note")
	}

	return nil
}

func share(c echo.Context) error {
	cc := c.(*CustomContext)
	s := cc.Server

	var dto core.ShareDTO
	if err := c.Bind(&dto); err != nil {
		return fmt.Errorf("invalid body format")
	}

	if err := isOwner(c, dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	user := s.Db.FindUser(dto.UserEmail)
	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "User with that email was not found",
		})
	}

	payload := map[string]string{
		"object":   fmt.Sprintf("note:%d", dto.NoteId),
		"relation": dto.Permission,
		"user":     fmt.Sprintf("user:%d", user.Id),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "error creating payload",
		})
	}

	err = s.SendAcl(jsonData)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to share the note",
		})
	}

	sw := core.SharedWithDTO{
		Email:      user.Email,
		Name:       user.Name,
		Permission: dto.Permission,
	}

	return c.JSON(http.StatusOK, sw)
}

func unshare(c echo.Context) error {
	cc := c.(*CustomContext)
	s := cc.Server

	var dto core.ShareDTO
	if err := c.Bind(&dto); err != nil {
		return fmt.Errorf("invalid body format")
	}

	if err := isOwner(c, dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	fmt.Println(dto.UserEmail)
	user := s.Db.FindUser(dto.UserEmail)
	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "User with that email was not found",
		})
	}

	payload := map[string]string{
		"object":   fmt.Sprintf("note:%d", dto.NoteId),
		"relation": dto.Permission,
		"user":     fmt.Sprintf("user:%d", user.Id),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "error creating payload",
		})
	}

	err = s.DeleteAcl(jsonData)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to delete the permissions",
		})
	}

	return c.NoContent(200)
}

func login(c echo.Context) error {
	cc := c.(*CustomContext)
	s := cc.Server

	var dto core.LoginDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	user, err := s.Db.LoginUser(dto)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "wrong credentials",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "wrong credentials",
		})
	}

	token := core.JwtGenerator(*user, "secretkey")

	return c.JSON(http.StatusOK, map[string]string{
		"token":  token,
		"userId": strconv.Itoa(user.Id),
		"email":  user.Email,
	})
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

func getShared(c echo.Context) error {
	cc := c.(*CustomContext)
	s := cc.Server

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	userId := strconv.FormatFloat(id, 'f', -1, 64)

	noteId, err := strconv.Atoi(c.Param("noteId"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid path parameter type",
		})
	}

	if b, err := s.CheckAcl(strconv.Itoa(noteId), "owner", userId); err != nil && !b {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "You are not the owner of this note",
		})
	}

	users, err := s.Db.GetAllUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	var ret []*core.SharedWithDTO
	permissions := []string{"editor", "viewer"}
	for _, user := range users {
		if user.Id == int(id) {
			continue
		}
		for _, p := range permissions {
			if b, err := s.CheckAcl(strconv.Itoa(noteId), p, strconv.Itoa(user.Id)); err == nil && b {
				sw := core.SharedWithDTO{
					Email:      user.Email,
					Name:       user.Name,
					Permission: p,
				}
				ret = append(ret, &sw)
				break
			}
		}
	}

	if ret == nil {
		return c.JSON(http.StatusOK, make([]string, 0))
	}

	return c.JSON(http.StatusOK, ret)
}
