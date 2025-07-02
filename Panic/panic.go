package main

func employee(name *string, age int) {
	if(age > 60) {
		panic("Age can not be greater than retirement age")
	}
	// else{
	// 	fmt.Println("Name: ", *name)
	// 	fmt.Println("Age: ", age)
	// }
}
func main() {
	name := "Nikhil Codes"
	age := 61
	employee(&name, age)
}