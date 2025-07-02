package main
import ("fmt")

type Person struct{
	firstName string
	lastName string
}

func (p Person) fullName() string {   // type of method is Person 
	return p.firstName + " " + p.lastName
}

func main() {
	person := Person{
		firstName: "Nikhil",
		lastName: "Yadav",
	}

	name := person.fullName()
	fmt.Println(`Name: `, name)
}