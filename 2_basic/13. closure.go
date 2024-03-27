package goBasic

import "fmt"

func plus(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func PlayWithClosure() {
	result := plus(1)(2)
	fmt.Printf("Closure: %v", result)
}