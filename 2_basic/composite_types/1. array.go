package goCompositeTypes

import (
	"fmt"
	"reflect"
)

/*
	Array has a fixed size and same types of items
*/

func PlayWithArray() {
	fmt.Println("------ Array with no element ---")
	var nilArr [3]string;
	fmt.Println(nilArr)

	fmt.Println("------ Array after changing item ---")
	names := [5]string{"mario", "puzo", "james", "yoshi", "kenta"}
	names[3] = "maria"
	fmt.Println(names, len(names))
	fmt.Printf("Type of names is a %v\n", reflect.TypeOf(names).Kind())
}