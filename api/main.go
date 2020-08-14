package main

import (
	"github.com/ono5/study-hacker/api/controller"
	"github.com/ono5/study-hacker/api/server"
)

func main() {
	server := server.NewServer()
	server.Get("/", controller.HelloHandler)
	server.Run()
}
