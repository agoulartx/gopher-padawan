package data_structures

import "fmt"

func Showarray() {
	var x [43]int

	for i:=0; i<43;i++{
		x[i] = (i)
	}
	fmt.Println(x)
}