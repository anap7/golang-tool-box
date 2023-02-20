package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

/*
Defer é uma clausula, que adia a execução de um determinado 
pedaço de código (função, função de algum pacote)
*/
func main() {
	//Será executada a função2 e depois a função alunoEstaAprovado 
	defer funcao1()
	funcao2()
	fmt.Println(alunoEstaAprovado(7, 8))
}