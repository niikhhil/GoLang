package main
import ("fmt")

func main() {
	h := struct{
		firstName string
		lastName string
	}{
		firstName: "Nikhil",
		lastName: "Yadav",
	}

	fmt.Println(h.firstName, h.lastName)
}