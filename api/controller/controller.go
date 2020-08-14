package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// HelloHandler is controller for hello feature
func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World2")
}
