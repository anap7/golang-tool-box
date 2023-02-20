package main

import "fmt"

func main() {
	//Declaração explícita: Deixando evitende o tipo da variável
	var var01 string = "first"
	
	//Declaração implícita: Deixando o tipo da variável oculto
	var02 := "second hiding type"

	fmt.Println(var01, var02)

	//Declaração em conjunto - Modo 1
	var (
		var05 string = "test-var05"
		var06 string = "test-var06"
	)

	fmt.Println(var05, var06)

	//Declaração em conjunto - Modo 2
	var07, var08 := "test-var07", "test-var08"

	fmt.Println(var07, var08)

	//Constante 
	const const01 string = "Valor constante: test-const01"

	const (
		const02 string = "test-const02"
		const03 string = "test-const03"
	)

	fmt.Println(const01, const02, const03)
}