package goDataTypes

import (
	"fmt"
	"sort"
)

/*
	A slice is a dynamically-sized array which is built on top of array (fixed-sizes)
*/

func PlayWithSlice() {
	// Slices: no fixed size
	var scores = []int{20, 40, 80}
	scores[2] = 60
	scores = append(scores, 100)
	sort.Ints(scores)

	fmt.Println(scores, len(scores))

	// Slices ranges
	names := [5]string{"mario", "puzo", "james", "yoshi", "kenta"}
	rangeOne := names[1: 3]
	rangeTwo := names[2:]
	rangeThree := names[:2]

	fmt.Println(rangeOne, len(rangeOne), cap(rangeOne))
	fmt.Println(rangeTwo, len(rangeTwo), cap(rangeTwo))
	fmt.Println(rangeThree, len(rangeThree), cap(rangeThree))
}