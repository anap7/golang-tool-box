package main

import "fmt"

/*O Worker Pool é um padrão utilizando quando existir uma fila de tarefas
grande para ser executada e fazer com que tenha vários processos pegando
cada item dessa fila e ir executando de forma independepente*/

func main() {
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		result := <- results
		fmt.Println(result)
	}
}

//Declarando que o canal tasks recebe dados apenas e o results envia dados
func worker(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		results <- fibonacci(number)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}