package main

import "fmt"

// canales con buffer
func main() {
	ca := make(chan int, 2)

	ca <- 46
	ca <- 45
	fmt.Println(<-ca)
	fmt.Println(<-ca)

	ca <- 44
	fmt.Println(<-ca)

	ca <- 39
	ca <- 38
	fmt.Println(<-ca)
	fmt.Println(<-ca)
}
