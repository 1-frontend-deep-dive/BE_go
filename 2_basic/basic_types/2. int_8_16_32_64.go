package goBasicTypes

import "fmt"

func PlayWithNumber() {
	// Ints
	var initOne int = 28
	var initTwo = 30
	initThree := 32
	fmt.Println(initOne, initTwo, initThree)

	// Signed integers
	// int8:  -128 to 127
	// int16: -32768 to 32767
	// int32: -2147483648 to 2147483647
	// int64: -9223372036854775808 to 9223372036854775807
	var signedIntOne int8 = 25
	var signedIntTwo int8 = -128
	fmt.Println(signedIntOne, signedIntTwo)

	// Unsigned integers
	// uint8:  0 to 255
	// uint16: 0 to 65535
	// uint32: 0 to 4294967295
	// uint64: 0 to 18446744073709551615
	var unSignedIntOne uint8 = 255
	fmt.Println(unSignedIntOne)

	// Floats
	// float32: -3.4e+38 to 3.4e+38
	// float64: -1.7e+308 to +1.7e+308
	var floatOne float32 = 123.78
	fmt.Println(floatOne)

	// Types conversion
	var pureInt8 int8 = 100
	var pureUint8 uint8 = uint8(pureInt8)

	fmt.Println(pureUint8)

}