package functions

import "fmt"


func Expression() {
	mytype := func(){
		fmt.Println("I'm a function expression")
	}
	mytype()

	fmt.Printf("mytype is: %T \n", mytype)
}
