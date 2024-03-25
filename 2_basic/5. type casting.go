package goBasic

import "fmt"

func PlayWithTypeCasting() {
	var i int = 42
	var f float32 = float32(i)
	var u uint = uint(f)

	fmt.Println(u)
}