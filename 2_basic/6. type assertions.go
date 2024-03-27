package goBasic

import (
	"fmt"
	"reflect"
)

func PlayWithTypeAssertions() {
	var i interface{} = "hello"
	fmt.Printf("Type i: %v - value: %v \n", reflect.TypeOf(i).Kind(), i)

	s, ok := i.(string) // value is hello
	fmt.Printf("Type s: %v - value: %v - isOk: %v\n", reflect.TypeOf(s).Kind(), s, ok)

	f1, ok := i.(float32) // value is 0
	fmt.Printf("Type f: %v - value: %v - isOk: %v\n", reflect.TypeOf(f1).Kind(), f1, ok)

	// f2 := i.(float32) // panic
	// fmt.Printf("Type f: %v - value: %v",  reflect.TypeOf(f2).Kind(), f2)

	switch i.(type) {
	case int:
		fmt.Println("I has int type")
	case string: 
		fmt.Println("I has int type")	
	default:
		fmt.Println("I is an unknown type")
	}
}