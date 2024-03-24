package goBasic

import "fmt"

// Go is a pass-by-value language
// Non-pointer Values: strings, ints, floats, booleans, arrays, structs
// Pointer Wrapper Values: slices, maps, functions


func passByReference(x map[string]string, keyName string, newValue string) {
	// the key of the map stores the pointer of the value
	// a new copy of people variable (x) was created along with the pointers of value
	x[keyName] = newValue
}

func passByValue(x string, newValue string) {
	x = newValue
}

func PlayWithPassByValue() {
	name := "mario";
	passByValue(name, "puzo");
	fmt.Println(name) // still mario

	fmt.Println("---------------------------------------")

	people := map[string]string {
		"name": "mario",
		"age": "27",
	}
	passByReference(people, "name", "puzo")
	fmt.Println(people) // changed to puzo
}