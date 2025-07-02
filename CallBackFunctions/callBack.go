package main

import (
	"fmt"
)

func addName(name string, callback func(string)) {
	callback(name)
}

func main() {
	
	addName("Nicks code", func(nm string) {
		fmt.Printf("Hi my name is %v", nm) 
	})
}