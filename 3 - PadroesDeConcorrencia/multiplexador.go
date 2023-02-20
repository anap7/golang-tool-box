package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*O Multiplexador é um padrão cuja ideia é pegar dois ou mais canais e juntar
em um único canal

A função escrever será chamada mais de uma vez e juntar esses canais e centrar a comunicação
em um canal só*/
func main() {
	channel := multiplex(write("Hello World!"), write("Hey Listen!"))

	for i:=0; i < 20; i++ {
		fmt.Println(<-channel)
		fmt.Println("===============================")
	}
}

//Dois canais que só recebem valores
func multiplex(channelInput1, channelInput2 <-chan string) <-chan string {
	channelOutput := make(chan string)

	go func() {
		for {
			select {
			//Caso exista uma mensagem disponível no canal, colocaremos no canal de saída
			case message := <-channelInput1:
				channelOutput <- message
			case message := <-channelInput2:
				channelOutput <- message
			}
		}
	}()

	return channelOutput
}

//A função recebe uma string e devolve um canal de string, um canal de uma direção só
func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido %s", text)
			//Pausa em tempo aleatório
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}