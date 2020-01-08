package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan bool)

	go enviar(par, impar, salir)
	recibir(par, impar, salir)
}

func enviar(p, i chan<- int, s chan<- bool) {
	for x := 0; x < 100; x++ {
		if x%2 == 0 {
			p <- x
		} else {
			i <- x
		}
	}
	close(s)
}

func recibir(p, i <-chan int, s <-chan bool) {
	for {
		select {
		case v := <-p:
			fmt.Println("Desde el canal Par:\t", v)
		case v := <-i:
			fmt.Println("Desde el canal Impar:\t", v)
		case i, ok := <-s:
			if !ok {
				fmt.Println("Desde como ok1:\t", i)
				return
			} else {
				fmt.Println("Desde como ok2:\t", i)
			}
		}
	}
}
