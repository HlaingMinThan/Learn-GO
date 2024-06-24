package main

import (
	"fmt"
	"strings"
)

func main() {
	welcome := "greeting"
	// fmt.Println(strings.Contains(welcome, "g"))
	fmt.Println(strings.ReplaceAll(welcome, "e", "a")) //graating
	fmt.Println(strings.Index(welcome, "ting"))        //4
	fmt.Println(strings.ToUpper(welcome))              //GREETING
	fmt.Println(strings.Split(welcome, "t"))           //GREETING
}
