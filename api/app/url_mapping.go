package app

import (
	"github.com/ono5/study-hacker/api/controllers"
	"github.com/ono5/study-hacker/api/server"
)

// URLMapping - Mapping Handler and URL
// Explorer is an nterface, so we can put a restraction on method such as GET, POST.
// because they should be assigned in Explorer interface!
func URLMapping(s server.Explorer) {
	s.GET("/", controllers.DBHandler)
}
