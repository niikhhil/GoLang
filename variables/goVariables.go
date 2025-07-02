package main

import ("fmt")

const PI = 3.14 // can also be declared in block like block variables

var name = "Nikhil Yadav" //using var keyword outside the main

func main () {
	var fruit = "Mango" // using var keyword inside the main        inferred type automatically

	number := 1  // type is inferred automatically by the compiler

	var animal string = "Lion"  // using := punisherar  string = "Lion" // assigning types to variable

	//Multiple declaration of variables
	 var one, two, three, four, five int = 1, 2, 3, 4, 5

	 // Declaring variables in a block	
	 var (
		a int
		b = 2
		c = "Hello world"
	 )

	fmt.Println( "Name:", name )
	fmt.Println( "Fruit:", fruit ) 
	fmt.Println( "Number:", number )
	fmt.Println( "Animal:", animal )
	fmt.Println("Multiple Variables:", one, two, three, four, five)
	fmt.Println( "Block Variables:", a, b, c)
	fmt.Println( "Contant:", PI)

	fmt.Println(len(animal))

}