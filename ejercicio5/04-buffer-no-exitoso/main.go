package main

import "fmt"

// canales con buffer - caso no exitoso (el buffer se llena)
func main() {
	ca := make(chan int, 1)

	ca <- 46
	ca <- 42

	fmt.Println(<-ca)
}
