package closures

import "fmt"

func wrapper() func () int {
	x := 40
	return func() int{
		x++
		return x
	}
}

func Closuremain()  {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())

}
