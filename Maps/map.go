package main

import ( "fmt" )

func main() {
	userData := map[string]int{
		"Nikhil": 20,
		"Mannat": 19,
	}
	fmt.Println(userData)

	for prop, value := range userData{
		fmt.Println(prop, value)
	}
}