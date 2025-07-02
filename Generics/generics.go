package main
import ("fmt")

// generic function 
func  printItem[T any](a , b T)  (T, T) {
	return a, b
}

func main() {
	num1, num2 := printItem(4, 8)
	num3, num4 := printItem(true, false)
	num5, num6 := printItem("Nikhil", "Yadav")
	fmt.Println(num1, num2)
	fmt.Println(num3, num4)
	fmt.Println(num5, num6)
}