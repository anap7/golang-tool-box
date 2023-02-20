
package main

import "fmt"

/*Closure: Funções que referenciam variáveis fora do seu corpo
Veja que a função closure retorna uma função
*/
func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}