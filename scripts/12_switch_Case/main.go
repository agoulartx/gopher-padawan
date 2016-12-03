package main

import "fmt"

func main(){
	// Single evaluation
	nome := "mario"
	switch nome {
	case "jhon":
		fmt.Println("Welcome Jhon")
	case "mario":
		fmt.Println("Welcome Mario")
	default:
		fmt.Println("Nenhum")

	}

	// Multiple evaluation
	name := "maria"
	switch name {
	case "jhon","carlos":
		fmt.Println("Welcome Jhon")
	case "rosa","maria":
		fmt.Println("Welcome Maria")
	default:
		fmt.Println("Nenhum")
	}
}