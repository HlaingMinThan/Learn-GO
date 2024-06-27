package main

import "fmt"

func main() {
	createBill()
	bill := newBill("HMT")
	bill.updateTip(10)
	bill.addItem("Onion", 5)
	fmt.Println(bill.format()) //not override the original name
}
