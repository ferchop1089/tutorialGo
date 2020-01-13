// Package mymath probando el comando doc de go
package mymath

// Sum suma los valores dados por un slice
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}
