package main

import (
	"fmt"
	"time"
)

/*Channels: Canal de comunicação, onde você pode enviar ou receber dados e usar para sincronizar nossas goroutines:*/
func main() {
	/*Para criar um canal usamos o método "make", dentro inserimos a palavra chave "chan" para criar o canal
	e como segundo parâmetro o tipo desse canal, no caso o tipo string.Esse canal pode enviar e receber somente
	strings*/
	channel := make(chan string)

	fmt.Println("Depois da função escrever começar a ser executada")

	//Criando uma goroutine
	go write("Hello World", channel)

	for {
		//(recebendo o valor) nessa posição da seta, significa que o canal aguarda o recebimento do valor
		message, open := <-channel
		/*E para evitar deadlock, verificamos se o canal está fechado e encerramos o processo
		de looping*/
		if !open {
			break
		}
		fmt.Println("Sintaxe 01: No loop do canal, segue a mensagem:", message)
	}

	//Mesmo exemplo acima, porém com a sintaxe mais simples
	for message := range channel {
		fmt.Println("Sintaxe 02: No loop do canal, segue a mensagem:", message)
	}

	fmt.Println("Fim do programa")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		//(enviando o valor) texto sendo inserido no canal(variavel <- valor == variavel == valor)
		channel <- text
		fmt.Println(text)
		time.Sleep(time.Second)
	}

	//Fechando o canal após o looping
	close(channel)
}