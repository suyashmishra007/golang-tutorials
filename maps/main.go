package main

import "fmt"

func main() {

	// name <-> grade
	studentGrade := make(map[string]int)
	studentGrade["Price"] = 10
	studentGrade["Bob"] = 40
	fmt.Println(studentGrade);
	studentGrade["Bob"] = 70
	studentGrade["Alice"] = 30
	fmt.Println(studentGrade);
	// delete(studentGrade,"Bob");
	fmt.Println(studentGrade);

	// Checking if the key exits or not.
	grades,exits := studentGrade["Alice"];

	fmt.Println("Grades of Alice",grades);
	fmt.Println("Bob exits",exits);

	for key,value := range studentGrade {
		fmt.Printf("Key is %s and marks is %d\n",key,value);
	}

	// Map can be initialized using the make function, or using a map literal.
	studentMarks := map[string]int{
		"Alice" : 10,
		"Bob": 20,
		"Charlie" : 30,
	}
	fmt.Println(studentMarks)
}