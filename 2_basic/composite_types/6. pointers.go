package goCompositeTypes

import "fmt"

/*
	A pointer holds the memory address of a value
*/

func updateName(name *string) {
	*name = "Iris"
}

func PlayWithPointer() {
	name := "Stephan"

	fmt.Println("Pointer address: \t", &name)
	fmt.Println("Before update name: \t", name)

	updateName(&name)
	
	fmt.Println("After update name: \t", name)
}