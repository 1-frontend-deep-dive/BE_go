package goCompositeTypes

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", cases.Title(language.Und, cases.NoLower).String(n))
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func plus(a int8, b int8) int8 {
	return a + b
}

func getFirstAndLastName(name string)(string, string) {
	names := strings.Split(name, " ")
	return names[0], names[1]
}

func PlayWithFunction() {
	sayGreeting("Mario")

	fmt.Println("---------------------------------------")
	cycleNames([]string {"cloud", "barret", "ben"}, sayGreeting)

	fmt.Println("---------------------------------------")
	fmt.Printf("Sum of %v and %v is %v \n", 2, 4, plus(2, 4))

	fmt.Println("---------------------------------------")
	firstName, lastName := getFirstAndLastName("James Cameron");
	fmt.Println("%v - %v \t", firstName, lastName)
}