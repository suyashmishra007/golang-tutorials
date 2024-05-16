package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	fmt.Println("Enter your Name");
// 	var name string;

// 	fmt.Scan(&name);
// 	fmt.Println("Name is :",name);
// }

func main() {
	fmt.Println("Enter your Name");
	
	reader := bufio.NewReader(os.Stdin);
	name,_ := reader.ReadString('\n')
	
	fmt.Println("Name is :",name);
}