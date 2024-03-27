package goBasic

import "fmt"

/*
 * Defer will be execute before the function end even the function throws errors
 * List of defer task are on stack, which means Last In First Out
 */

func PlayWithDefer() {
	defer fmt.Println("Hello")
	fmt.Println("World")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}