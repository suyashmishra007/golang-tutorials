package main

import "fmt"

func main() {
	fmt.Println("Starting of the program");
	defer fmt.Println("Middle of the program");
	defer fmt.Println("After Middle of the program");
	fmt.Println("End of the program");

}