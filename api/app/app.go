package app

import (
	"github.com/ono5/study-hacker/api/server"
)

// SetupServer - Set up server to run echo engine
func SetupServer() server.Server {
	server := server.NewServer()
	UrlMapping(server)
	server.Run()
	return server
}

// Run - application start
func Run() {
	SetupServer()
}
