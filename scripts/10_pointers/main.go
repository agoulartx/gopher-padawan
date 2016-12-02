package main

import "fmt"

func zero(x *int) {
	*x = 0
}
func main(){
	a := 42
	// Print a and it's memory address
	fmt.Println(a)
	fmt.Println(&a)

	// Start a pointer to a's memory address
	var b *int = &a

	// Print pointer b and the value it's stores
	fmt.Println(b)
	fmt.Println(*b)

	// Change the value of pointer b, and also changes the value of a due to the pointer
	*b = 24
	fmt.Println(*b)
	fmt.Println(a)

	// -------------------------------
	// Using pointer to change a variable value on the global scope


	// define x
	fmt.Println("---------")
	x := 5
	fmt.Println(x)
	zero(&x)
	fmt.Println(x)
}