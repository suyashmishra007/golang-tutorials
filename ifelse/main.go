package main

import "fmt"

func main() {
	x := 10
	if x >= 10 {
		fmt.Println("x is greater or equal than 10");
	}else{
		fmt.Println("x is smaller than 0");
	}
	
	
	z := 0
	if z > 0 {
		fmt.Println("z is positive");
	}else if z == 0 {
		fmt.Println("z is zero");
	}else {
		fmt.Println("z is negative");
	}
}