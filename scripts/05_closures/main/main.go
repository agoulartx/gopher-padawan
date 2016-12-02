package main

import (
	"github.com/agoulartx/gopher-padawan/scripts/05_closures/innerscope_closure"
	"fmt"
)

func main() {
	fmt.Println("---------------")
	fmt.Println("Anonymous func output")
	closures.Anonymousmain()
	fmt.Println("---------------")

	fmt.Println("Inner scope output")
	closures.Innermain()
	fmt.Println("---------------")

	fmt.Println("Closures output")
	closures.Closuremain()

}
