package goCompositeTypes

import (
	"fmt"
	"reflect"
)

/*
 * A slice is a descriptor of an array segment
 * It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).
 */

func PlayWithSlice() {
	fmt.Println("------ slice with no fixed-sized ---")
	var scores = []int{20, 40, 80} // Slice with no fixed-sizes
	fmt.Println(scores, len(scores), cap(scores))
	fmt.Printf("Type of scores is a %v\n", reflect.TypeOf(scores).Kind())


	fmt.Println("------ slice from an array ---")
	names := [5]string{"mario", "puzo", "james", "yoshi", "kenta"}
	rangeOne := names[1: 3] // Slice is created from array
	rangeOne[1] = "brent" // james in names array will be changed to brent

	fmt.Println(names, len(names), cap(names))
	fmt.Println(rangeOne, len(rangeOne), cap(rangeOne))
	fmt.Printf("Is rangeOne a slice: %v\n", reflect.TypeOf(rangeOne).Kind() == reflect.Slice)
}