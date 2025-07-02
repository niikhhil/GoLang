// Slices are used to store multiple values of same type in a single variable, however unlike arrays, the length of the slices can grow and shrink as you see fit

package main
import (
	"fmt"
)

func main () {
	mySlice := []int{1, 2, 3, 4, 5} // no need to give the size as required in arrays
	fmt.Println(mySlice[0:3]) // last index non inclusive 

	// creating slice using make method
	slice := make([]string, 3, 5) // take 3 arguments type, length and capacity
	fmt.Println("Length : ", len(slice))
	fmt.Println("Capacity : ", cap(slice))
	fmt.Println("slice", slice)

	// append in slice  append(slice, values)
	 mySlice = append(mySlice, 6, 7, 8, 9)

	 fmt.Println("Updated slice: ",mySlice)

	 ppls := []string{"Nikhil", "Ravi", "Jubaida", "Salmaan"}

	 for index, value := range ppls {
		fmt.Println(index, value)
	 }
}