package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(2000 * time.Millisecond) // Simulationg some work
	fmt.Println("Function Ended Here")
}

func sayHi() {
	fmt.Println("Hi Hi Hi")
}
func main() {
	go sayHello()
	sayHi()

	// Wait for moment to allow the goroutines to finish
	time.Sleep(2500 * time.Millisecond)
}
