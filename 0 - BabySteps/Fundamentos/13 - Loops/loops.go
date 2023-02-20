package main

import (
	"fmt"
	"time"
)

func doWhile(num int) {
	i := 0;

	for i < num {
		i++ 
		fmt.Println("Do While: Incrementando ...", i)
		//Delay por um segundo antes printar no próximo loop
		time.Sleep(time.Second)
	}
}

func forLoop(num int) {
	for i := 0; i < num; i++ {
		fmt.Println("For: Incrementando ...", i)
		time.Sleep(time.Second)
	}
}

func readAList() {
	names := [4]string{"Gui", "Guizinho", "Luizinho", "Zezinho"}
	
	/*
	*	Index: Posicição do array
	* name: item no array
	* range names é a lista que você quer interar
	*/

	//Maneira 1
	for index, name := range names {
		fmt.Println(index, name)
	}

	//Maneira 2
	for _, name := range names {
		fmt.Println(name)
	}

	//Com string
	for index, letter := range "Word" {
		fmt.Println(index, letter)
	}

	//Com Map
	user := map[string]string {
		"name": "Paulinha",
		"age": "12",
	}

	for key, values := range user {
		fmt.Println(key, values)
	}

	//Com struct não é possível fazer loop
}

func main() {
	readAList()
	//doWhile(5)
	//forLoop(5)
}