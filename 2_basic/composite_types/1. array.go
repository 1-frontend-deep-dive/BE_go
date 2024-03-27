package goCompositeTypes

import "fmt"

/*
	Array has a fixed size and same types of items
*/

func PlayWithArray() {
	var nilArr [3]string;
	fmt.Println(nilArr)

	names := [5]string{"mario", "puzo", "james", "yoshi", "kenta"}
	names[3] = "maria"
	fmt.Println(names, len(names))
}