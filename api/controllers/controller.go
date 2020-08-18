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

// CreateHandler is controller for create feature
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

// UpdateHandler is controller for update feature
func UpdateHandler(c echo.Context) error {
	data := new(models.Language)
	err := c.Bind(data)
	utils.LogFatal("Cannot paramter from context: %v\n", err)
	m := driver.ConnectDB("study", "language")
	mr := repository.NewMongoRepo(m, data)
	data, err = mr.Update(data)
	if err != nil {
		return c.String(
			http.StatusInternalServerError,
			fmt.Sprintf("%v", err),
		)
	}
	return c.JSON(http.StatusOK, data)
}

// DeleteHandler is controller for date feature
func DeleteHandler(c echo.Context) error {
	id := c.Param("id")

	data := new(models.Language)
	m := driver.ConnectDB("study", "language")
	mr := repository.NewMongoRepo(m, data)
	id, err := mr.Delete(id)
	if err != nil {
		return c.String(
			http.StatusInternalServerError,
			fmt.Sprintf("%v", err),
		)
	}
	return c.String(http.StatusOK, id)
}
