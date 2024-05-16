package main

import "fmt"

func simpleFunction() {
	fmt.Println("this is simple function");
}

func add(a, b int) int {
	return a + b;
}

func multiply(a,b int) (result int) {
	result = a * b;
	return
}
func main() {
	ans := multiply(2,2);
	fmt.Println("Ans is" ,ans);
}