package goFoundation

import "fmt"

func PlayWithCondition() {
	age := 28

	if (age < 10) {
		fmt.Println("Age is less than 10")
	} else if age < 20 {
		fmt.Println("Age is less than 20")
	} else {
		fmt.Println("Age is more than 20")
	}
}