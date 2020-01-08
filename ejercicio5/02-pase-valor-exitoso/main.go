package main

import "fmt"

// canales sin buffer
func main() {
	ca := make(chan int)

	go func() {
		ca <- 46
	}()

	fmt.Println(<-ca)
}
