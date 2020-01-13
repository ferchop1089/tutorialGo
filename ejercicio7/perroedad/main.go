package main

import (
	"fmt"
	"tutorialGo/ejercicio7/perroedad/perro"
)

type canino struct {
	nombre string
	edad   int
}

func main() {
	c1 := canino{"Luna", perro.Edad(2)}
	fmt.Println(c1)
}
