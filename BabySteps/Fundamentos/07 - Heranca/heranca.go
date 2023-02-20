package main

import (
	"fmt"
)

type person struct {
	name string
	lastName string
	age uint8
	hight uint8
}

//Para realizar a herança, basta colocar o nome do struct
type student struct {
	person
	course string
	college string
}

func main() {
	fmt.Println("Herança")

	p1 := person{"Paulinho", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := student{p1, "Engenharia", "USP"}
	fmt.Println(e1.hight)
}