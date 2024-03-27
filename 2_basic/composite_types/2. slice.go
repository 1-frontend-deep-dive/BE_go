package goCompositeTypes

import (
	"fmt"
)

/*
 * A slice is a descriptor of an array segment
 * It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).
 */

func PlayWithSlice() {
	var scores = []int{20, 40, 80} // Slice with no fixed-sizes
	fmt.Println(scores, len(scores), cap(scores))


	names := [5]string{"mario", "puzo", "james", "yoshi", "kenta"}
	rangeOne := names[1: 3] // Slice is created from array
	rangeOne[1] = "brent" // james in names array will be changed to brent

	fmt.Println(names, len(names), cap(names))
	fmt.Println(rangeOne, len(rangeOne), cap(rangeOne))
}