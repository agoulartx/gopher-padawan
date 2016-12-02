package main

import "fmt"

// unitary declaration
const p = "yoda"

// multiple declaration

const (
	size  = 34
	value = 54.6
	color = "red"
)

func main(){

	fmt.Println("p - ", p)
	fmt.Println("size - ", size)
	fmt.Println("value - ", value)
	fmt.Println("color - ", color)

}
