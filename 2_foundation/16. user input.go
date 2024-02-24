package goFoundation

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, error := r.ReadString('\n')

	return strings.TrimSpace(input), error
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Print("Created the bill - ", b.name)

	return b
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