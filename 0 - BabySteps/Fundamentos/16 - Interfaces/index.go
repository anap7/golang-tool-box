package main

import (
	"fmt"
	"math"
)

type forma interface {
	//Indicando que essa interface tem um método chamado área que retorna um float64
	area() float64
}

type retangulo struct {
	altura  float64
	largura  float64
}

type circulo struct {
	raio float64
}

/*O Go implementa a interface de maneira implícita
o r retangulo é passado como método */
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f", f.area())
	fmt.Printf("\n")
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)


	c := circulo{10}
	escreverArea(c)
}