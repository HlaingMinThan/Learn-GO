package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"name": "hlaingminthan",
		"age":  "20",
	}

	fmt.Println(person)
	person["name"] = "updated"
	person["contact"] = "contact"
	person["age"] = "updated 1w"
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person)
}
