package main

import (
	"fmt"
)

func testPassByValue(n *string) {
	*n = "override"
}

func main() {
	name := "Hlaingminthan"
	testPassByValue(&name)
	fmt.Println(name) //not override the original name
}
