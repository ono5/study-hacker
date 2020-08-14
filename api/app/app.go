package app

import (
	"github.com/ono5/study-hacker/api/server"
)

func SetupServer() server.Server {
	server := server.NewServer()
	UrlMapping(server)
	server.Run()
	return server
}

func Run() {
	SetupServer()
}
