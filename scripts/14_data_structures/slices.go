package data_structures

import "fmt"

func Showslice() {
	mySlice := []int{1,2,3,4,5,6,7,8,9,0}
	fmt.Println(mySlice)

	// Slicing

	fmt.Println(mySlice[2:5])

	// Access by index

	fmt.Println(mySlice[0])
	fmt.Println("mySlice"[4])

}