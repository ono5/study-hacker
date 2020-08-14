package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ono5/study-hacker/api/driver"
	"github.com/ono5/study-hacker/api/models"
	"github.com/ono5/study-hacker/api/repository"
)

// DBHandler is controller for hello feature
func DBHandler(c echo.Context) error {
	m := driver.ConnectDB("study", "language")
	data := &models.Language{
		Japanese: "こんにちは",
		English:  "Hello",
	}
	mr := repository.NewMongoRepo(m, data)
	id, err := mr.Create()

	if err != nil {
		return c.String(
			http.StatusInternalServerError,
			fmt.Sprintf("%v", err))
	}

	return c.String(http.StatusOK, id)
}
