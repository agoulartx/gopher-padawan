package main

import "fmt"

func main() {

	// Shorthand declaration

	a := 10
	b := "golang"
	c := 4.55
	d := true

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Println("-----------\n")

	// Print the variable's type

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)

	// Declare zero values

	fmt.Println("-----------\n")
	var e int
	var f string
	var g float64
	var h bool

	fmt.Printf("Float 64: %v \n", e)
	fmt.Printf("String: %v \n", f)
	fmt.Printf("Float: %v \n", g)
	fmt.Printf("Bool: %v \n", h)


}
