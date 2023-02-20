package main

import "fmt"

func recuperarExecucao() {
	/*O recover é uma clausula utilizada quando o panic é
	acionado, é como se fosse um catch e vc consegue seguir
	com o fluxo, veja que na linha 15 chamamos o defer, pq
	queremos chamar um processo antes e ver se dá problema,
	caso dê o recover vai ser acionado em seguida*/
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	/*A panic interrompe o fluxo normal do programa
	O erro é algo que você pode capturar e seguir ocm o 
	fluxo do progrma, já o panic interrompe todo o processo */
	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução!")
}