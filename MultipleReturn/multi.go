package main

import("fmt"
		"strconv"
)

func swap(a, b int) (int, int) {
	return b, a
}

// returning result and error together 
func convert(str string) (int, error) {
	result, err := strconv.Atoi(str)
	return result, err
}
func main() {
	
	a, b := swap(4, 5)
	fmt.Println(a, b)
	
	num, err := convert("1234")
	if(err == nil) {
		fmt.Println("String converted to integer: ",num)
	}else{
		fmt.Println("Error:", err)
	}
}