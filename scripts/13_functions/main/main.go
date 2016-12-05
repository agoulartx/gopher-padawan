package main

import (
	"github.com/agoulartx/gopher-padawan/scripts/13_functions"
	"fmt"
)

func main(){
	// Arguments
	functions.Singleprint("Hello!")
	functions.Multiprint("Welcome,","Jhon")
	fmt.Println("-------------------")

	// Returns
	a := functions.Greet("Augusto","Goulart")
	fmt.Println(a)
	fmt.Println(functions.Namedgreet("Jhon","Doe"))
	fmt.Println(functions.Doublereturn("Carlos","Mario"))
	fmt.Println("-------------------")

	// Variadic
	v := functions.Variadicsum(188.4,487.6,8.9,222.1,543.6)
	fmt.Println(v)
	fmt.Println("-------------------")

	// Func Expression

	functions.Expression()
	fmt.Println("-------------------")

	// Closures

	closure := functions.Wrapper()
	//Wrapper is a function that returns a function
	fmt.Println("X value: ",closure())
	fmt.Println("X value: ",closure())
	fmt.Println("X value: ",closure())
	fmt.Println("X value: ",closure())

	// Callback
	fmt.Println("-------------------")
	functions.Printslice([]int{1,2,3,4,5,4,3,2,1}, func(n int){
		fmt.Println(n)
	})
	fmt.Println("-------------------")

	xs := functions.Filter([]int{1,2,3,4,5,6}, func(n int) bool{
		return n % 2 ==0
	})
	fmt.Println(xs)


}
