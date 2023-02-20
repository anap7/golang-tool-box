package main

import (
	"fmt"
	"reflect"
)
func main() {
	fmt.Println("Arrays e Slices")

	// Introdução - Arrays e Slices
	var array1 [5]string

	//Populando um array
	array1[0] = "Posição 1"
	fmt.Println(array1)

	//Declarando um array e populando
	array2 := [5]string{"Step 1", "Step 2", "Step 3", "Step 4", "Step 5"}
	fmt.Println(array2)

	/*Declarando um array, porém aplicando um destructuring e passando os valores
	Isso representa uma forma dinâmica de popular um array*/
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//Criando um slice
	/*A diferença é a sua flexibilidade de tamanhos
		- Um slice não é um array

	*/

	//
	slice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(slice[1:4])
	//Retorna o tipo da variável, utilizamos o reflect
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	//Append -> Adicionar um item novo ao slice e rtorna um novo slice com o novo item
	slice = append(slice, 200)
	fmt.Println(slice)

	//"Pegando um pedaço" dos itens do slice de outro array
	slice2 := array2[2:4]
	fmt.Println(array2)
	fmt.Println(slice2)

	//Array Internos
	/*
	Vamos utilizar o make para criar um slice,
	alocando um espaço da memória para uma determinada
	variável ou lista
	*/
	fmt.Println("------------------------")
	//Cria um array de 15 posições e retorna as 10 primeiras posições
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Tamanho ocupado (10)
	fmt.Println(cap(slice3)) //Capacidade, total do array mesmo com espaços vazios (15)
	
	//Inserindo um item ao slice, ocupando a posição 11
	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Tamanho ocupado (11)
	fmt.Println(cap(slice3)) //Capacidade, total do array mesmo com espaços vazios (15)
}