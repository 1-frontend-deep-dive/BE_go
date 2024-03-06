package goFoundation

import "fmt"

func PlayWithChannel() {
	// Create a channel to send and receive data
	c := make(chan int)

	// Launch a gorountine to send a 42 through channel
	go func() {
		c <- 42
	}()

	// Receive the sent 42 through channel
	fmt.Println(<-c)
}