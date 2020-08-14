package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ono5/study-hacker/api/driver"
	"github.com/ono5/study-hacker/api/models"
	"github.com/ono5/study-hacker/api/repository"
	"github.com/ono5/study-hacker/api/utils"
)

// CreateHandler is controller for hello feature
func CreateHandler(c echo.Context) error {
	data := new(models.Language)
	err := c.Bind(data)
	utils.LogFatal("Cannot paramter from context: %v\n", err)
	m := driver.ConnectDB("study", "language")
	mr := repository.NewMongoRepo(m, data)
	id, err := mr.Create()

	if err != nil {
		return c.String(
			http.StatusInternalServerError,
			fmt.Sprintf("%v", err))
	}

	return c.String(http.StatusOK, id)
}
