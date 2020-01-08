package main

import "fmt"

func main() {
	fmt.Println("Llamando a método preparar")
	preparar()
	fmt.Println("Terminó la prueba de panic")
}

func preparar() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Se lanzó una excepción:", r)
		}
	}()
	fmt.Println("Llamando a método lanzarPanic")
	lanzarPanic(1)
	fmt.Println("Éste texto nunca es mostrado en consola dado que en lanzarPanic se lanza una excepción")
}

func lanzarPanic(i int) {
	if i > 3 {
		panic(fmt.Sprintf("i es mayor a 3, actualmente es %v", i))
	}
	fmt.Println("Valor de i:", i)
	lanzarPanic(i + 1)
}
