package main

import "fmt"

type user struct {
	name string
	age int8
	address address
}

type address struct {
	logradouro string
	numero int8
}

func main() {
	fmt.Println("Struct -> Coleção de campos, com um nome e tipo")

	//Realizando a instância do usuário
	var user1 user 
	user1.name = "Ana"
	user1.age = 23
	fmt.Println(user1)
	fmt.Println(user1.name,"\n",user1.age)

	adressExample := address{"Rua dos Bobos", 0}

	//Modo 2 - Realizando a instância do usuário
	user2 := user{"Pokemon", 34, adressExample} 
	fmt.Println(user2)
	fmt.Println(user2.address.logradouro)

	//Modo 3 - Realizando a instância do usuário, somente um atributo
	user3 := user{name: "Pokemon"} 
	fmt.Println(user3)

	user4 := user{age: 34} 
	fmt.Println(user4)

}