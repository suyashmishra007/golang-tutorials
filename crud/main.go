package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// Send GET request to the API endpoint
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	// Decode JSON response into a Todo struct
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the todo with explicit field labels
	fmt.Println("UserID:", todo.UserID)
	fmt.Println("ID:", todo.Id)
	fmt.Println("Title:", todo.Title)
	fmt.Println("Completed:", todo.Completed)
}
