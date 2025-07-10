package main

import("fmt")

func main() {

	// single defer statement executes at the end of the execution of surrounding elements

	// multiple defer function execute in LIFO order 
	defer fmt.Println("First Function")
	defer fmt.Println("Second Function")
	defer fmt.Println("Third Function")
}