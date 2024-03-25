package goBasic

import "fmt"

func PlayWithLoop() {
	fmt.Println("--------------- Loop with for ---------------")

	x := 0
	for x < 5 {
		fmt.Println("Value of x is: ", x)
		x++
	}

	fmt.Println("--------------- Loop with for in array ---------------")

	names := [5]string{"mario", "puzzo", "james", "yoshi", "kenta"}
	for i := 0; i < len(names); i++ {
		fmt.Println("Item of names array is: ", names[i])
	}

	fmt.Println("--------------- Loop with range in array ---------------")

	for index, value := range names {
		fmt.Printf("Item of names array is %v with index %v \n", value, index)
	}

	fmt.Println("--------------- Loop with range in map ---------------")

	animals := map[string]string{
		"dog": "Spike",
		"cat": "Tom",
		"mouse": "Jerry",
	}

	for aniKey, aniValue := range animals {
		fmt.Printf("The %s has name %s\n", aniKey, aniValue)
	}
}