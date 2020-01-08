package main

import "fmt"

// Canal sin buffer - ejemplo de bloqueo (Se requiere de una gorutina que envie y otra que reciba)
func main() {
	ca := make(chan int)
	ca <- 42
	fmt.Println(<-ca)
}