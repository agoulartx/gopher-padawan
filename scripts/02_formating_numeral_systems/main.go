package main

import "fmt"

func main() {
	// %d -> Decimal
	// %b -> Binary
	// %x -> Hexa lowercase
	// %X -> Hera uppercase
	// %q -> utf-8 table

	// Prints all the conversions above between 0-100
	for i := 0; i <= 100; i++ {
		fmt.Printf("Decimal: %d | Binary: %b | HexaL: %x | HexaU: %X | UTF-8: %q \n", i,i,i,i,i)
	}
}
