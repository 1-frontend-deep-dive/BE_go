package goDeeper

import (
	"fmt"
	"time"
)

/*
 * A goroutine is a lightweight thread managed by the Go runtime and
 * executes concurrently with the rest of the program
 */

 func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func PlayWithGoRoutines() {
	go say("world")
	say("hello")
}