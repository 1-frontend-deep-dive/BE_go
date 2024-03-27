package goBasicTypes

import "fmt"

func PlayWithString() {
	// Println
	var firstName string = "Mario"
	var middleName = " "
	lastName := "Puzo" // shorthand of above, can only used in a function
	
	fullName := fmt.Sprintf("%s - %s - %s", firstName, middleName, lastName)
	fmt.Println(fullName)

	// Printf
	person1 := "James"
	person2 := "Alisa"

	fmt.Printf("We're %v and %v \n", person1, person2)
	fmt.Printf("My age is %q and name is %q \n", 24, person2)
	fmt.Printf("Age is of type %T \n", 24)
	fmt.Printf("You scored %f points \n", 225.44788)
	fmt.Printf("You scored %0.1f points \n", 225.44788)
}