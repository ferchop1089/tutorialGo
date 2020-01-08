package main

import "fmt"

func main() {
	ca := make(chan int)

	go func() {
		for x := 1; x < 5; x++ {
			ca <- x
		}
		close(ca)
	}()

	for x := range ca  {
		fmt.Println("x: ",x)
	}
}
