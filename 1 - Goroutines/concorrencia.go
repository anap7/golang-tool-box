package main

import (
	"fmt"
	"time"
)

func main() {
	/*O "go" significa que você quer que seja executada
	a função na linha 14, porém não espera que essa função
	termine o processo para que seja chamada a próxima função
	na linha 15 e sim que seja chamada de maneira revezada
	ambas as funções */
	go write("Hey")
	write("Hey Listen!")
}

func write(text string) {
	for {
		fmt.Println(text)
		//Um pequeno delay para ver pausadamente a impressão do texto
		time.Sleep(time.Second)
	}
}