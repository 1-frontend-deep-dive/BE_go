package goFoundation

import "fmt"

// A pointer holds the memory address of a value

func updateName(x *string) {
	*x = "wedge"
}

func PlayWithPointer() {
	name := "tifa"


	namePointer := &name
	namePointerValue := *namePointer

	fmt.Println("Memory address of name variable\t:", namePointer)
	fmt.Println("Value of name pointer \t\t:", namePointerValue)

	updateName(namePointer)
	fmt.Println("New value of name variable\t:", name)
}