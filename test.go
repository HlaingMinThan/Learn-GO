package main

import (
	"fmt"
)

func testPassByReference(m map[string]string) {
	m["age"] = "override"
}
func testPassByValue(n string) {
	n = "override"
}

func main() {
	//pass by reference -> map,slices,functions
	person := map[string]string{
		"name": "hlaingminthan",
		"age":  "20",
	}
	testPassByReference(person)
	fmt.Println(person) //override the original person map data

	//pass by value -> int,floats,strings,struct,array,bools
	name := "Hlaingminthan"
	testPassByValue(name)
	fmt.Println(name) //not override the original name
}
