package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Infinite Loop
	counter := 0
	for {
		fmt.Println("Infinite Loop", counter);
		counter++;
		if counter == 10 {
			break;
		}
	}

	// Range keyword
	numbers := [5]int {21,32,43,14,25}
	for index,value := range(numbers){
		fmt.Println(index,value);
	}

	strs := "Hello World";
	for index,char := range strs{
		fmt.Printf("%d and %c \n",index,char)
	}
}