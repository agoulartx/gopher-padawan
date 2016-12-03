package functions

import "fmt"

func Variadicsum(values ...float64) string {
	fmt.Printf("Valus: %v \n", values)
	fmt.Printf("Values type: %T \n",values)
	var total float64
	for _, v := range values {
		total += v
	}
	return fmt.Sprint("Total: ",total)
}
