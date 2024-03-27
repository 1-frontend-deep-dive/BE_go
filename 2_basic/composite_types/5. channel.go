package goCompositeTypes

import (
	"fmt"
	"math/rand"
)

/*
 * Channels are the pipes that connect concurrent goroutines
 * You can send values into channels from one goroutine and receive those values into another goroutine
 */

func getRandomNum(poolNum int) int {
	randomNum := rand.Intn(poolNum)
	return randomNum
}

func calculateValue(chanInt chan int) {
	randomNum := getRandomNum(100)
	chanInt <- randomNum
}

func PlayWithChannel() {
	chanInt := make(chan int)
	defer close(chanInt) // close thing that needs to be closed

	go 	calculateValue(chanInt)

	result := <- chanInt;
	fmt.Println("Random number from channel: ", result)
}