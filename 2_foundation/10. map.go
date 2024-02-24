package goFoundation

import "fmt"

/*
	Maps are used to store data values in key:value pairs
*/

func PlayWithMap() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          2.99,
		"coffee": 3.55,
	}

	fmt.Println(menu["pie"])

	fmt.Println("---------------------------------------")

	// loop in maps
	for key, value := range menu {
		fmt.Printf("%v\t - %v\n", key, value)
	}

	fmt.Println("---------------------------------------")

	// ints as key
	phoneBook := map[int]string {
		1111: "mario",
		2222: "puzo",
		3333: "james",
		4444: "yoshi",
	}

	phoneBook[4444] = "kenta"
	delete(phoneBook, 2222)

	for key, value := range phoneBook {
		fmt.Printf("%v\t - %v\n", key, value)
	}
}