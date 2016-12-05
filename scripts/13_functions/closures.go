package functions

func Wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}