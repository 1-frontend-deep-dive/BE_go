package goBasic

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, error := r.ReadString('\n')

	return strings.TrimSpace(input), error
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	bill, cashier := newBill(name)
	fmt.Printf("The bill %s is created by cashier %s", bill.name, cashier.fullname)

	return bill
}

func promptOptions (b Bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option \n- a (Add item)\n- s (Save bill)\n- t (Add tip)\n", reader)
	fmt.Println(opt)
}

func PlayWithUserInput() {
	myBill := createBill()
	promptOptions(myBill)
}