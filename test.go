package main

import "fmt" //formatting

// go will automatically fire this "main" function
// How to run - go run filename
func main() {
	// age := 23
	// name := "hlaing min than"
	number := 23.55

	//Printf (print with formatted string)
	// %v
	// fmt.Printf("My name is %v.my age is %v", name, age)

	// %q
	// fmt.Printf("My name is %q.my age is %q", name, age) //quote is only for string

	// %T
	// fmt.Printf("name is a type of %T", name)

	// %f
	// %0.1f
	// fmt.Printf("age is %0.1f", number)

	//Sprintf (save print with formatted string)
	var saved = fmt.Sprintf("age is %0.1f", number)
	fmt.Println(saved)
}
