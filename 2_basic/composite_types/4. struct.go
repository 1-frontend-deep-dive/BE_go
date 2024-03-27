package goCompositeTypes

import "fmt"

/*
 Struct is a user-defined type to store a collection of different fields into a single field
*/

type Cashier struct {
	fullname string
	age int
}

type Bill struct {
	name string
	items map[string]float64
	tip  float64
}

func newBill(name string) (Bill, Cashier) {
	bill := Bill{
		name:  name,
		items: map[string]float64{
			"coffee": 2.99,
		},
		tip:   0,
	}

	cashier := Cashier {
		"Marina",
		28,
	}

	return bill, cashier
}

func PlayWithStruct() {
	myBill, cashier := newBill("mario's bill")
	
	fmt.Println("My bill: ", myBill)
	fmt.Println("Cashier: ", cashier)
}