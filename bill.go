package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n') //this will run when user hit enter
	return input, err
}

func createBill() bill {
	//get user input string value
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Type Your Name", reader)
	name = strings.TrimSpace(name)
	bill := newBill(name)
	fmt.Println("created bill for " + bill.name)
	return bill
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput("a - add item, s - save bill, t - tip bill", reader)
	option = strings.TrimSpace(option)

	switch option {
	case "a":
		name, _ := getInput("Item Name:", reader)
		price, _ := getInput("Item Price:", reader)
		name = strings.TrimSpace(name)
		price = strings.TrimSpace(price)
		p, _ := strconv.ParseFloat(price, 64)
		b.addItem(name, p)
		promptOptions(b)
	case "s":
		b.save()
	case "t":
		tip, _ := getInput("Tip Amount:", reader)
		tip = strings.TrimSpace(tip)
		t, _ := strconv.ParseFloat(tip, 64)
		b.updateTip(t)
		fmt.Println("added tip to bill")
		promptOptions(b)
	default:
		fmt.Println("Invalid Option")
		promptOptions(b)
	}
}

func newBill(n string) bill {
	b := bill{
		name:  n,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b bill) format() string {
	fs := "Bill breakdown\n"
	var total float64 = 0
	total += b.tip
	for k, v := range b.items {
		total += v
		fs += fmt.Sprintf("%-15v  ....$%.2f\n", k+":", v)
	}

	fs += fmt.Sprintf("%-15v  ....$%.2f\n", "Tip:", b.tip)
	fs += fmt.Sprintf("%-15v  ....$%.2f\n", "Total:", total)
	return fs
}

func (b *bill) updateTip(tip float64) {
	// we don't need asterisk on struct b, go willl automatically dereference
	b.tip = tip
}
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
func (b *bill) save() {
	bytes := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", bytes, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}
