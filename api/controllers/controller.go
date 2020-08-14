package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ono5/study-hacker/api/driver"
	"github.com/ono5/study-hacker/api/models"
	"github.com/ono5/study-hacker/api/repository"
)

// HelloHandler is controller for hello feature
func HelloHandler(c echo.Context) error {
	m := driver.ConnectDB()
	m.GetCollection("study", "language")
	id, err := repository.Create(m, &models.Language{
		Japanese: "こんにちは!",
		English:  "Hello!",
	})

	if err != nil {
		return c.String(
			http.StatusInternalServerError,
			fmt.Sprintf("%v", err))
	}

	return c.String(http.StatusOK, id)
}
