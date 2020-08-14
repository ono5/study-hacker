package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ono5/study-hacker/api/server"
)

func main() {
	server := server.NewServer()
	server.Get("/", hello)
	server.Run()
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
