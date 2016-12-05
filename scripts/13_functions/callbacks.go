package functions


func Printslice(numbers []int, callback func(int)) {
	for _, n := range  numbers {
		callback(n)
	}
}