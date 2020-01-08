package main

import "fmt"

func main() {
	ca := make(chan int, 1)
	ca <- 21
	fmt.Println(<-ca)
	fmt.Printf("%T\n",ca)
}
