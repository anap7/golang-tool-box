package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*WaitGroup é usado para esperar que todas as goroutines
	inseridas na função terminem juntas. Isso de maneira sincronizada
	
	2 rotinas vão iniciar de maneira revezada, porém devem ser finalizadas
	juntas, por isso o uso do WaitGroup

	Nota: se um WaitGroup for explicitamente passado para funções, 
	isso deve ser feito por ponteiro */
	var waitGroup sync.WaitGroup

	/*Aqui eu informo que existem 2 goroutines que precisam
	ser inseridas nessa fila de waitGroup e que precisa
	esperar para que ambas finalizem*/
	waitGroup.Add(2)

	//Declarando uma função anônima que é uma goroutine
	/*
		lembrando que ambas vão rodar de maneira revezada
	*/
	go func() {
		write("Olá Mundo")
		//Ao chegar no Done() o waitGroup.Add removerá um item da fila
		waitGroup.Done()
	}()
	
	go func() {
		write("Programando em GO")
		waitGroup.Done()
	}()

	/*Esperar a contagem das goroutines 
	chegarem em 0, inicialmente declaramos 2 
	
	*/
	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}