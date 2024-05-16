package main

import "fmt"

func main() {
	// Declare and initialize a slice
	numbers := []int {1,2,3,4};
	
	// append more elements
	numbers = append(numbers, 5,6,7,8,9,10)

	// Access and print slice elements
	fmt.Println(numbers);

	// Length of the slice
	fmt.Println(len(numbers))

	// empty slice
	name := []int{}
	fmt.Println(name);

	// make function
	nums := make([]int,3,5); // 3 is current length , 5 is capacity
	fmt.Println("Slice", nums)
	fmt.Println("len", len(nums))
	fmt.Println("cap", len(nums))
	
}