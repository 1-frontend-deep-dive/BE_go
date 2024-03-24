package goDataTypes

import "fmt"

/*
	Array has fixed size and same types of items
*/

func PlayWithArray() {
	var nilArr [3]string;
	var age = [3]int{20, 25, 30} // lint error if there is the fourth item
	names := [5]string{"mario", "puzo", "james", "yoshi", "kenta"}
	names[3] = "maria"

	fmt.Println(nilArr)
	fmt.Println(age, len(age))
	fmt.Println(names, len(names))
}