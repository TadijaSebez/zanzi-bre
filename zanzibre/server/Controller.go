package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syndtr/goleveldb/leveldb"
)

func dbPutTest(c echo.Context) error {
	db, err := leveldb.OpenFile("zanzibase", nil)

	if err != nil {
		return err
	}

	defer db.Close()

	err = db.Put([]byte("key"), []byte("value"), nil)

	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Successful put")
}

func dbGetTest(c echo.Context) error {
	db, err := leveldb.OpenFile("zanzibase", nil)

	if err != nil {
		return err
	}

	defer db.Close()

	val, err := db.Get([]byte("key"), nil)

	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(val))
}
