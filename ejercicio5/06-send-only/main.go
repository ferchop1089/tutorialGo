package main

import "fmt"

// canal para sólo enviar información
func main() {
	ca := make(chan <- int, 1)
	ca <- 21
	//fmt.Println(<-ca) // genera error ya que el canal sólo envía
	fmt.Printf("%T\n",ca)
}
