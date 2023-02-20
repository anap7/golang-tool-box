package main

import "fmt"

//Como o nome já diz, declaramos o nome da variável de retorno 
func calculosMatematicos(n1, n2 int) (soma int, subtracao int, divisao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	divisao = n1 / n2
	return
}

func main() {
	soma, subtracao , divisao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao, divisao)
}