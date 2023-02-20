package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

/*
* Nessa função, veja que estamos declarando o tipo somente
para o n2 e o n1 não tem a declaração do tipo, no entanto, o
go é fortemente tipado e por isso o n1 vai pegar o tipo
declarado do n2, então n1 será do tipo int8 e espera
2 tipos de retorno, colocado em parenteses (int8 int8)
*/
func calc(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	result := somar(12, 34)
	fmt.Println(result)

	//Variável tipo função
	var f = func() {
		fmt.Println("Nova função -> Chamando a função de soma")

		result := somar(2, 2)
		fmt.Println("Resutado da função de soma", result)
	}

	//Variável tipo função tipo string que retorna uma string
	var ftxt = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultString := ftxt("Olá como vai")
	fmt.Println("Chamando uma função de string pela variável:", resultString)

	f()

	//Pegando dois retornos
	resultSoma, resultSub := calc(20, 14)
	fmt.Println("Resultado do calculo de soma:", resultSoma, "\nResultado do calculo de sub:", resultSub)

	//Pegando somente o retorno da subtração
	_, resultSubOnly := calc(75, 14)
	fmt.Println("Retorno do calculo somente de subtração:", resultSubOnly)

	//Pegando somente o retorno da subtração
	resultSomaOnly, _ := calc(35, 7)
	fmt.Println("Retorno do calculo somente de soma:", resultSomaOnly)



}