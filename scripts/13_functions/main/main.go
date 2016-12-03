package main

import (
	"github.com/agoulartx/gopher-padawan/scripts/13_functions"
	"fmt"
)

func main(){
	functions.Singleprint("Hello!")
	functions.Multiprint("Welcome,","Jhon")
	a := functions.Greet("Augusto","Goulart")
	fmt.Println(a)

	fmt.Println(functions.Namedgreet("Jhon","Doe"))
	fmt.Println(functions.Doublereturn("Carlos","Mario"))
}
