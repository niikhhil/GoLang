package main
import ("fmt")

func main() {
	x := 10
	addressOfx := &x  // can be written as var addressOfx *int = &x
	// * used to specify the variable of pointer type
	// & used to store the address of another variable
	//fmt.Println("Value of x: ", x)
	fmt.Println("Address of x: ", addressOfx)
	fmt.Println("Value at x: ", *addressOfx)
}