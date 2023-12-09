package dataTypes

import "fmt"

func PlayWithLoop() {
	x := 0

	for x < 5 {
		fmt.Println("Value of x is: ", x)
		x++
	}

	fmt.Println("---------------------------------------")
	names := [5]string{"mario", "puzzo", "james", "yoshi", "kenta"}

	for i := 0; i < len(names); i++ {
		fmt.Println("Item of names array is: ", names[i])
	}

	fmt.Println("---------------------------------------")

	for index, value := range names {
		fmt.Printf("Item of names array is %v with index %v \n", value, index)
	}

	fmt.Println("---------------------------------------")

	for _, value := range names {
		fmt.Printf("Item of names array is %v \n", value)
	}
}