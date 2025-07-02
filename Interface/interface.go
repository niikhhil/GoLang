package main

import "fmt"

type Animal interface {
	getInfo() string
}

type Cat struct {
	Name  string
	Color string
}

type Dog struct {
	Name string
	Breed string
}
func (c Cat) getInfo() string {
	return fmt.Sprintf("Name of the cat is %s and its color is %s", c.Name, c.Color)
}

func (d Dog) getInfo() string {
	return fmt.Sprintf("Name of Dog is %s and its breed is %s", d.Name, d.Breed)
}

func printAnimal (animal Animal) {
	fmt.Println(animal.getInfo())
}

func main() {

	cat := Cat {
		Name: "Abella",
		Color: "Orange",
	}
	printAnimal(cat)

	dog := Dog{
		Name: "Coco",
		Breed: "Golden Retriever",
	}
	printAnimal(dog)

}