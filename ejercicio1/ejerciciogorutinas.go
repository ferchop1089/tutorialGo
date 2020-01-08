package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("Número de CPUs: %v\n", runtime.NumCPU())

	wg.Add(2)

	llamada1()
	go llamada2()
	go func() {
		fmt.Println("Llamada 3")
		wg.Done()
	}()

	fmt.Printf("Número de gorutinas: %v\n", runtime.NumGoroutine())
	wg.Wait()
}

func llamada1()  {
	fmt.Println("Llamada 1")
}

func llamada2()  {
	fmt.Println("Llamada 2")
	wg.Done()
}
