package dataTypes

import (
	"fmt"
	"sort"
	"strings"
)

func PlayWithStdLib() {
	fmt.Println("---------- string ----------")

	greeting := "hello world"
	fmt.Println("Contains: \t", strings.Contains(greeting, "hello"))
	fmt.Println("ReplaceAll: \t", strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println("ToUpper: \t", strings.ToUpper(greeting))
	fmt.Println("Index: \t\t", strings.Index(greeting, "o"))
	fmt.Println("Split: \t\t", strings.Split(greeting, " "))

	fmt.Println("---------- sort ----------")

	names := [5]string{"mario", "puzo", "james", "yoshi", "kenta"}
	ages := []int{4, 89, 34, 56, 11, 53, 85, 30, 43}

	sort.Strings(names[:])
	sort.Ints(ages)
	indexOf34 := sort.SearchInts(ages, 34)

	fmt.Println("sort.Strings: \t\t", names)
	fmt.Println("sort.SearchStrings: \t", sort.SearchStrings(names[:], "puzo"))
	fmt.Println("sort.Ints: \t\t", ages)
	fmt.Println("sort.SearchInts: \t", indexOf34)
}