package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(`Strings in GO`)

	stringOne := `Nikhil yadav`
	stringTwo := `Nicks the coder`
	fmt.Println(`Length of String: `,len(stringOne)) //calculating length
	fmt.Println(strings.Compare(stringTwo, stringOne)) //.Compare(s1, s2)  result = s1 - s2   { 1 (s1 > s2), 0, -1 }
	fmt.Println(strings.ToLower(stringOne)) // to lower case
	fmt.Println(strings.ToUpper(stringOne)) // to upper case 
}