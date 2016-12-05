package data_structures

import "fmt"

func Showslice() {
	mySlice := []int{1,2,3,4,5,6,7,8,9,0}
	fmt.Println(mySlice)
	fmt.Println("-----------")

	// Slicing

	fmt.Println(mySlice[2:5])

	// Access by index
	fmt.Println("-----------")

	fmt.Println(mySlice[4])
	fmt.Println("-----------")

	newSlice := make([]int, 0, 5)

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println("-----------")

	for i := 0; i<80;i++{
		newSlice = append(newSlice, i)
		fmt.Println("Len:", len(newSlice), "Cap:",cap(newSlice), "Value:", newSlice[i])
	}


}