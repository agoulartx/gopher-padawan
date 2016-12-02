package main

import "fmt"

func main() {
	i := 0
	for i<=200 {
		if i%2 == 0 {
			fmt.Println("Par")
		}else {
			fmt.Println("Impar")
		}
		i++
	}
}
