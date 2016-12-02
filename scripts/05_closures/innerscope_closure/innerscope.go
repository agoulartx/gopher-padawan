package closures


import "fmt"

func Innermain() {
	x := 42
	fmt.Println(x)
	// Create an inner scope
	{
		fmt.Printf("%d from inner scope \n",x)
		y := "Do or do not, there is no try"
		fmt.Println(y)
	}
	// Error
	// fmt.Println(y) -> y no define
}
