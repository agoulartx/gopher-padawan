package closures

import "fmt"

func Anonymousmain() {
	x := 0
	// Declare an anonymous function
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}