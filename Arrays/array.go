package main
import (
	"fmt"
)
func main()  {
	var arr1 = [5]int{1,2,3,4,5}
	arr2 := [5]int{1,2,3,4,5}

	var empty [5]int  // empty array or zero array   empty := [5]int{}
	fmt.Println(`Empty Array ->`,empty)

	fmt.Println(arr1)
	fmt.Println(arr2)

	
	for i := 0; i < len(arr1); i++ {
		fmt.Print(arr1[i], " ")
	}
}