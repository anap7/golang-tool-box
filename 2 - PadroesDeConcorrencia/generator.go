package main

import (
	"fmt"
	"time"
)

/*O Generator é um padrão com uma função que vai encapsular uma goroutine
e devolver um canal*/
func main() {
	//A função retorna um canal que só recebe dados
	channel := write("Hello World")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

//A função recebe uma string e devolve um canal de string, um canal de uma direção só
func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}