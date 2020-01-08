package main

import "fmt"

// canal para sólo recibir información
func main() {
	ca := make(<- chan int, 1)
	// ca <- 21 // genera error ya que el canal sólo recibe
	fmt.Println(<-ca)
	fmt.Printf("%T\n",ca)
}
