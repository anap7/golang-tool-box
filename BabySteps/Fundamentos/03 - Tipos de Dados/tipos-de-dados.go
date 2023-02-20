package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int64 = 100000000000
	fmt.Println("int64 ->", number)

	//Aceita número negativo
	var number2 int32 = -1000000000
	fmt.Println("int32 negative number ->", number2)

	//Utilizando alias para o int32 - rune === int32
	var number3 rune = 34444
	fmt.Println("int32 === rune ->", number3)

	//Utilizando alias para o int8 - int8 === byte
	var number4 byte = 34
	fmt.Println("int8 === byte ->", number4)

	//Utilizando float
	var numberFloat float32 = 32.99
	fmt.Println("Float 32 ->", numberFloat)

	//Strings
	var str string = "Texto"
	fmt.Println("String ->", str)

	str2 := "Texto 2"
	fmt.Println("String 2 ->", str2)

	var bool1 bool = false
	fmt.Println("Boolean ->", bool1)

	var bool2 bool
	fmt.Println("Boolean ->", bool2)

	//errors é o pacote nativo de erros do golang
	//o tipo "error", muito utilizado
	var erro error = errors.New("Internal Error")
	fmt.Println("Error ->", erro)
}