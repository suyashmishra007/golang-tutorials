package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = 101; 
}
func main() {
	// Declare a pointer and a variable
	// var num int = 2
	// var ptr *int
	// ptr = &num
	// fmt.Println(ptr);


	num := "Suyash" 
	ptr := &num
	fmt.Println("pointer contains",ptr);
	fmt.Println("data contains through pointer",*ptr);

	// default is 'nil'
	var pointer *int ;
	fmt.Println(pointer);


	value := 10
	modifyValueByReference(&value);
	fmt.Println("Value is ", value);

}