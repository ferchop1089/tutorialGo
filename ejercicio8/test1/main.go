package main

import "fmt"

func main() {
	rta := Sum(1, 2, 3)
	fmt.Println("Resultado:", rta)
}

// Sum suma los valores dados al slice
func Sum(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}
