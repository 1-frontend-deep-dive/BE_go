package goBasic

import (
	"fmt"
)

type Cart struct {
	name string
	items map[string]float64
	tip  float64
}

func newCart(name string) Cart {
	b := Cart{
		name:  name,
		items: map[string]float64{
			"Coffee": 2.99,
			"Cake": 3.99,
		},
		tip: 0,
	}

	return b
}

// the format function can be used only in Card struct
func (c Cart) format() string {
	fs := "...::: Cart breakdown:::... \n"
	var total float64 = 0

	for key, value := range c.items {
		fs += fmt.Sprintf("%-30v ...$%v \n", key + ":", value)
		total += value
	}

	fs += fmt.Sprintf("%-30v ...$%v \n", "Tip", c.tip)
	fs += fmt.Sprintf("%-30v ...$%0.2f", "Total:", total)
	return fs
}

func (c *Cart) updateTip(tip float64) {
	c.tip = tip
}

func (c *Cart) addItem(name string, price float64) {
	c.items[name] = price
}

func PlayWithReceiverFunction() {
	myCard := newCart("mario's cart")
	myCard.addItem("Onion soup", 4.5)
	myCard.updateTip(10)
	
	fmt.Println(myCard.format())
}