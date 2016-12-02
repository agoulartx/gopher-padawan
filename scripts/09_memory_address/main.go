package main

import "fmt"

func main() {
	a := 43
	var idade int32

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Print("idade: ")
	fmt.Scan(&idade)
	fmt.Println(idade)
}
