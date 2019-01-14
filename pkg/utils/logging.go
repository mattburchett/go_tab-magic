package utils

import (
	"log"
	"time"
)

//LogElapsedTime provides a function that can be used to log the time elapsed between when this function is invoked, to the time the returned function is invoked.
//Adding "defer LogElapsedTime("my log")" at the beginning of any function will log the time it takes to execute that function.
func LogElapsedTime(what string) func() {
	start := time.Now()
	return func() {
		log.Printf("%s, Total Time Taken: %v\n", what, time.Since(start))
	}
}
