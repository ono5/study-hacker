package utils

import "log"

// LogFatal checks err
func LogFatal(msg string, err error) {
	if err != nil {
		log.Fatal(err)
	}
}
