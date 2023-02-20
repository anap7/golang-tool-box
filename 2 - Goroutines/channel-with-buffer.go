package main

import (
	"fmt"
)

/*Channels: Canal de comunicação, onde você pode enviar ou receber dados e usar para sincronizar nossas goroutines:*/
func main() {
	/*Para criar um canal usamos o método "make", dentro inserimos a palavra chave "chan" para criar o canal
	e como segundo parâmetro o tipo desse canal, no caso o tipo string.Esse canal pode enviar e receber somente
	strings*/
	/*
		O parâmetro "2", simboliza que estou criando um canal com buffer, ou seja, um canal com capacidade
		e com o limite, controlando a capacidade e o tráfego de dados evitando um deadlock

		EM RESUMO:
			Um canal com buffer só vai bloquear o tráfego de dados quando o canal atingir sua capacidade 
			máxima, já um canal sem buffer sempre vai bloquear quando existir a operação de envio e recebimento
	*/
	channel := make(chan string, 2)
	//Enviando o valor no canal
	channel <- "Hello World"
	channel <- "Coding in Golang"
	//Caso inserimos mais um dado no canal, teremos um deadlock porque ultrapassou o limite do canal
	//channel <- "Test"

	//Recebendo o valor no canal
	message := <- channel
	//Para receber a segunda mensagem, basta apontar para uma nova variável
	message2 := <- channel
	
	fmt.Println(message)
	fmt.Println(message2)

}
