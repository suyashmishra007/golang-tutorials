package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close();
	//We Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Println(sb)
}
