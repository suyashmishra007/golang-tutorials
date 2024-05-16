package main

import (
	"fmt"
	"net/url"
)

// https://glitchyhitchy.medium.com/handling-urls-in-golang-b59a2b26a202#:~:text=To%20parse%20a%20URL%20in,and%20returns%20a%20URL%20type.&text=The%20URL%20type%20in%20Golang,we%20cannot%20modify%20it%20directly.
func main() {
	// URL := "https://www.example.com/path/to/resource?param1=value1&param2=value2";
	URL := "https://www.example.com/path/to/resource?param1=value1&param2=value2#fragment"
	parsedUrl, _ := url.Parse(URL)

	fmt.Printf("Type of parsedUrl : %T", parsedUrl)

	u, _ := url.Parse(URL)

	fmt.Println("Protocol:", u.Scheme)
	fmt.Println("Hostname:", u.Hostname())
	fmt.Println("Path:", u.Path)
	fmt.Println("Query:", u.RawQuery)
	fmt.Println("Fragment:", u.Fragment)

	// The URL type in Golang is immutable, which means that we cannot modify it directly.

	// Modify the protocol
	u.Scheme = "http"

	// Modify the hostname
	u.Host = "localhost:8080"

	// Modify the path
	u.Path = "/new/path"

	// Add a new query parameter
	q := u.Query()
	q.Set("param3", "value3")
	u.RawQuery = q.Encode()

	fmt.Println(u.String())

}
