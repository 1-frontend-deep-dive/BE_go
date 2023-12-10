package goFoundation

import "fmt"

func PlayWithSlice() {
	// Slices: no fixed size
	var scores = []int{20, 40, 80}
	scores[2] = 60
	scores = append(scores, 100)

	fmt.Println(scores, len(scores))

	// Slices ranges
	names := [5]string{"mario", "puzo", "james", "yoshi", "kenta"}
	rangeOne := names[1: 3]
	rangeTwo := names[2:]
	rangeThree := names[:2]

	fmt.Println(rangeOne, len(rangeOne))
	fmt.Println(rangeTwo, len(rangeTwo))
	fmt.Println(rangeThree, len(rangeThree))
}