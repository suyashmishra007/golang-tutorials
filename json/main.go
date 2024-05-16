package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {

	person := Person{Name: "John", Age: 32, IsAdult: true}
	fmt.Println("person data : ", person)

	// convert person to json : Marshal
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	// decodoing json data
	var personData Person
	err = json.Unmarshal(jsonData, &personData)

	if err != nil {
		fmt.Println("Error Unmarshalling")
		return
	}
	fmt.Println("Person Data", personData);
}
