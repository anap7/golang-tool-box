package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estruturas de controle")

	num := -10

	if num > 15 {
		fmt.Println("O número é maior que 15")
	} else {
		fmt.Println("O número é menor que 15")
	}

	//If init: Um fenômeno do go de declarar uma variável dentro do if
	/*Por ser uma variável criada dentro do escopo do if, vc não consegue 
	utilizar essa variável fora do escopo*/
	if otherNumber := num; otherNumber > 0 {
		fmt.Println("O número é maior que 0")
	} else {
		fmt.Println("O número é menor que 0")
	}
}