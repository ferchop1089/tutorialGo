package main

import "fmt"

func main() {
	ca := make(chan int)

	go enviar(ca)
	recibir(ca)
}

func enviar(c chan <- int) {
	c <- 23
}

func recibir(c <- chan int){
	fmt.Println(<-c)
}