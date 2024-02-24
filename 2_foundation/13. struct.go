package goFoundation

import "fmt"

/*
 Struct is a user-defined type to store a collection of different fields into a single field
*/

type Bill struct {
	name string
	items map[string]float64
	tip  float64
}

func newBill(name string) Bill {
	b := Bill{
		name:  name,
		items: map[string]float64{
			"coffee": 2.99,
		},
		tip:   0,
	}

	return b
}

func PlayWithStruct() {
	myBill := newBill("mario's bill")
	fmt.Println(myBill)
}