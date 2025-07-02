package main
import ("fmt")

type Person struct{
	name string
	age int
	salary int
	job string
}
type Intro struct{
	firstName string
	lastName string
}
type Programmer struct{
	Intro 
	isProgrammer bool
}
func main() {
	 var personOne Person

	 personOne.name = "Nikhil Yadav"
	 personOne.age = 21
	 personOne.salary = 100000
	 personOne.job = "Software Engineer"
	 fmt.Println(personOne)
	 name := Intro{
	 	firstName: "Nikhil",
	 	lastName: "Yadav",
	 }
	 fmt.Println(name)


	// nested structs 
	h := Programmer{
		Intro :Intro{ // Intro act as a property here so equals to is not required 
			firstName: "Nikhl",
			lastName: "Yadav",
		},
		isProgrammer: false,
	}
	fmt.Println(h)

}