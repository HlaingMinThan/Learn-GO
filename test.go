package main

import (
	"fmt"
	"sort"
)

func main() {
	// welcome := "greeting"
	// fmt.Println(strings.Contains(welcome, "g"))
	// fmt.Println(strings.ReplaceAll(welcome, "e", "a")) //graating
	// fmt.Println(strings.Index(welcome, "ting"))        //4
	// fmt.Println(strings.ToUpper(welcome))              //GREETING
	// fmt.Println(strings.Split(welcome, "t"))           //GREETING

	//sort and search slices int
	// numbers := []int{5, 10, 3}
	// sort.Ints(numbers) //override the original slices
	// fmt.Println(numbers)

	// ages := []int{10, 20, 30, 40, 40, 50}
	// fmt.Println(sort.SearchInts(ages, 40))

	//sort and search slices string
	names := []string{"mario", "luigi", "yoshi"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "yoshi"))
}
