package main

import "fmt"

type Person struct {
	FistName string
	LastName string
	Age      int
}
type Contact struct {
	Email string
	Phone string
}
type Address struct {
	House string
	Area  string
	State string
}

type Employee struct {
	Person_Details  Person
	Contact_Details Contact
	Address_Details Address
}

func main() {
	var p1 Person
	// 1st method
	p1.Age = 22
	p1.FistName = "Suyash"
	p1.LastName = "Mishra"
	fmt.Println(p1)

	// 2nd method
	p2 := Person{
		FistName: "Akash",
		LastName: "Mittal",
		Age:      12,
	}
	fmt.Println(p2)

	// 3rd method : Return a pointer to the struct
	var p3 = new(Person)
	p3.Age = 21
	p3.FistName = "Rashmi"
	p3.LastName = "Mittal"
	fmt.Println(p3) // p3 , Acts like a pointer

	var e1 Employee
	e1.Person_Details = p1
	e1.Address_Details = Address{
		House: "32/1",
		Area:  "Kanpur",
		State: "Uttar Pradesh",
	}
	e1.Contact_Details = Contact{
		Email: "test@gmail.com",
		Phone: "213123132",
	}
	fmt.Println(e1.Contact_Details)
}
