package main

import (
	"fmt"
)

func main() {
	bill := newBill("HMT")

	fmt.Println(bill.format()) //not override the original name
}
