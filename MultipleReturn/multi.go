package main

import("fmt")

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	
	a, b := swap(4, 5)
	fmt.Println(a, b)
	
}