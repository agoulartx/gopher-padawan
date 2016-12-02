package main

import "fmt"



func main() {

	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

}
