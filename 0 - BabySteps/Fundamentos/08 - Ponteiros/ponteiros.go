package main

import (
	"fmt"
)
func main() {
	fmt.Println("Herança")

	var variavel1 int = 10
	var variavel2 int = variavel1

	//Ambas serão 10
	fmt.Println(variavel1, variavel2)

	//Acrescenda + 1, ficando 11
	variavel2++
	fmt.Println(variavel1, variavel2)

	var variavel3 int
	//Criando um ponteiro de int - guarda o endereço de memória de um valor inteiro
	var ponteiro *int

	variavel3 = 100
	//Jogando uma variável para o endereço de memória, apenas usando o "&"
	ponteiro = &variavel3

	//Será retornado: 100 e o endereço de memória onde o 100 está armazeado
	fmt.Println(variavel3, ponteiro)

	//Vendo o valor que está salvo no endereço de memória (desreferenciação)
	variavel3 = 150
	//Será retornado: 150 e 150, porque estou vendo o valor armazenado no endereço de memória através do "*"
	fmt.Println(variavel3, *ponteiro)
}