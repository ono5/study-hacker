package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Explorer is a bunch of explorer methods
type Explorer interface {
	GET(path string, fn echo.HandlerFunc)
	POST(path string, fn echo.HandlerFunc)
}

// Server stands for Echo Server
type Server struct {
	engine *echo.Echo
}

// NewServer is a constoructor for Server
func NewServer() *Server {
	s := Server{engine: echo.New()}
	// SetUp Logger
	s.engine.Use(middleware.Logger())
	// SetUp Recover
	s.engine.Use(middleware.Recover())
	return &s
}

// GET is GET Method to register handler function
func (s *Server) GET(path string, fn echo.HandlerFunc) {
	s.engine.GET(path, fn)
}

// POST is POST Method to register handler function
func (s *Server) POST(path string, fn echo.HandlerFunc) {
	s.engine.POST(path, fn)
}

// Run starts echo server
func (s *Server) Run() {
	s.engine.Logger.Fatal(s.engine.Start(":8080"))
}
