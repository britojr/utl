package errchk

import "log"

// Check validates error and prints a log message
func Check(err error, message string) {
	if err != nil {
		log.Printf("%v: err(%v)\n", message, err)
		panic(err)
	}
}
