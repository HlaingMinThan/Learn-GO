package main

import "fmt"

func main() {

	//while loop example
	// i := 0
	// for i < 10 {
	// 	fmt.Printf("i is %v \n", i)
	// 	i++
	// }

	//for loop example
	names := []string{"hmt", "test"}
	// for j := 0; j < len(names); j++ {
	// 	fmt.Printf("name %v is %v \n", j+1, names[j])
	// }

	//for range loop example
	for index, name := range names {
		fmt.Printf("name %v is %v \n", index+1, name)
	}
}
