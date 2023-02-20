package main

import (
	"fmt"
)

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	//Veja que recebe qualquer valor
	generica("string")
	generica(1)
	generica(true)

}