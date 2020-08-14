package app

import (
	"github.com/ono5/study-hacker/api/server"
)

// SetupServer - Set up server to run echo engine
func SetupServer() {
	s := server.NewServer()
	URLMapping(s)
	s.Run()
}

// Run - application start
func Run() {
	SetupServer()
}
