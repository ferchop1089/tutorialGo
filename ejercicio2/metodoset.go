package main

import "fmt"

type humano interface {
	hablar()
}

type persona struct {
	nombres string
	edad int
}

func main() {
	p := persona{
		nombres: "Eduard",
		edad:    30,
	}
	diAlgo1(p)
	diAlgo2(&p)
	diAlgo3(&p)
}

func (p *persona) hablar()  {
	fmt.Printf("%s de %d años está hablando\n", p.nombres,p.edad)
}

func diAlgo1(p persona) {
	p.hablar()
}

func diAlgo2(p *persona) {
	p.hablar()
}

func diAlgo3(p humano) {
	p.hablar()
}