package app

import (
	"github.com/ono5/study-hacker/api/controller"
	"github.com/ono5/study-hacker/api/server"
)

func UrlMapping(s server.Server) {
	s.GET("/", controller.HelloHandler)
	s.GET("/hello", controller.HelloHandler)
}
