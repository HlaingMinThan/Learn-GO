package main

import "fmt"

//formatting

func main() {
	//array
	// var ages [3]int = [3]int{10, 20}

	// var ages = [3]int{1, 2, 3}
	// ages[1] = 100
	// ages := [3]int{1, 2, 3}

	// fmt.Println(ages)
	// fmt.Println(len(ages)) //get array length

	//slices push -> apppend can only used on slices
	slices := []int{1, 2, 3, 4}
	slices = append(slices, 30)

	// fmt.Println(slices)
	// fmt.Println(len(slices))
	// fmt.Println(slices[1:3]) // ranges[index:index] -> ans : [2,3]
	fmt.Println(slices[:3]) // ranges[index:index] -> ans : [1,2,3]
	fmt.Println(slices[1:]) // ranges[index:index] -> ans : [2,3,4,30]
}
