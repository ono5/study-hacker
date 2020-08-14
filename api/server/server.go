package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Exploer interface {
	Get(path string, fn echo.HandlerFunc)
}

type Server struct {
	engin *echo.Echo
}

func NewServer() Server {
	s := Server{engin: echo.New()}
	// SetUp Logger
	s.engin.Use(middleware.Logger())
	// SetUp Recover
	s.engin.Use(middleware.Recover())
	return s
}

func (s *Server) Get(path string, fn echo.HandlerFunc) {
	s.engin.GET(path, fn)
}
func (s *Server) Run() {
	s.engin.Logger.Fatal(s.engin.Start(":8080"))
}
