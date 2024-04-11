package goDeeper

import "fmt"

/*
 * With generics, we don't need to create multi function for different types with same logic
 */

func findIndex[T comparable](s []T, x T) int {
	for key, value := range s {
		if value == x {
			return key
		}
	}

	return -1
}

func getSpeed[SpeedType float32 | float64](s SpeedType) float64 {
	return float64(s)
}

type Car[T any] struct {
	name string
	speed T
}

type HumanHeight interface {
	float32 | int
}

func PlayWithGenerics() {
	fmt.Println("-------- Generics in func 1 --------")

	index1 := findIndex([]int{1, 7, 4, 8, 3, 8, 2}, 8)
	fmt.Println("Found number 5 at index:", index1)

	index2 := findIndex([]string{"dog", "cat", "fish"}, "fish")
	fmt.Println("Found fish at index:", index2)

	fmt.Println("-------- Generics in func 2 --------")

	speed := getSpeed(100.5)
	fmt.Printf("The vehicle is running at %v km/h\n", speed)

	fmt.Println("-------- Generics in struct --------")

	car := Car[float32]{
		name: "BMW",
		speed: 100.5,
	}
	fmt.Println(car)

	fmt.Println("-------- Generics in interface --------")
	
	// const humanHeight HumanHeight = 187
	// fmt.Printf("This person has %v as height\n", humanHeight)
}