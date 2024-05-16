package main

import "fmt"

func main() {
	fmt.Println("Learning Arrays")

	var name[5] string;
	name[0] = "0rice";
	name[1] = "1rice";
	name[2] = "2rice";
	name[3] = "3rice";
	name[4] = "4rice";

	var numbers = [5] int {11,12,13,14,15}

	// Length of array
	fmt.Println("Lenght of array numbers is ", len(numbers));
	
	// Accessing specific element
	fmt.Println(numbers[1]);

}