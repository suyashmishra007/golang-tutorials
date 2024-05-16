package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,mango,bananna"
	parts := strings.Split(data,",") // spliting on char
	fmt.Println(parts)
	
	str := "one one teo three four five six"
	cnt := strings.Count(str,"one")
	fmt.Println(cnt)

	str = "   Hello, Go    "
	fmt.Println(strings.TrimSpace(str)) // Removing extra space from front and back

	// Concat
	str1 := "Suyash"
	str2 := "Mishra"
	fmt.Println(strings.Join([]string{str1,str2}," "))
}