package main

import "fmt"

//Aguarda vários números como parâmetros (espera de 1 a n "ints")
func soma(numeros ...int) int {
	total := 0
	/*
	O underline no for _ indica quando vc não 
	quer utilizar o index, somente o valor do slice/array
	*/
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 0)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo", 10, 2, 3, 4, 5, 6, 7)
}