package main

import "fmt"

var n int

/*
A função init sempre será chamada primeira e depois a main
e você pode declará-la por arquivo
*/
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}