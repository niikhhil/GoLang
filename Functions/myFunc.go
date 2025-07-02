package main
import ("fmt")

// variadic functions
func showNumbers(numbers ...int) int { //unlimited inputs
	sum := 0
	for i := range numbers {
		sum += i
	}
	return sum
}

// simple function
func printName(Name string) {
	fmt.Println("Hello " + Name)
}

func add(x int, y int) int{
	return x + y
}

// Function that returns a closure
func incrementer() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}

func main() {
	printName("Nikhil")
	sum := showNumbers(1,2,3,4,5,6,7,8,9)
	fmt.Println("Addition of 5 and 10 is :", add(5, 10))
	fmt.Println("Sum of numbers is", sum)

	//anonymous function
	greet := func (name string){
		fmt.Println("Hello " + name +  ", this is a anonymous function")
	}
	greet("Nicks")

	 incr := incrementer()

    fmt.Println(incr()) // Output: 1
    fmt.Println(incr()) // Output: 2
}