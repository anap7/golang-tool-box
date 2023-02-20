package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")


	user := map[string]string {
		"name": "Paulinha",
		"lastName": "Silva",
	}

	fmt.Println(user["name"])
	fmt.Println(user["lastName"])

	/*Aqui eu indico que eu espero um map que recebe
	outro map, ou seja, um atributo que espera chaves e
	atributos dentro {} */
	user2 := map[string]map[string]string {
		"name": {
			"fisrt": "Ana",
			"last": "Flavia",
		},
		"birthDate": {
			"day": "14",
			"month": "12",
			"year": "1998",
		},
		"hobs": {
			"tt": "test",
			"tt2": "test2",
		},
	}

	fmt.Println(user2["name"]["fisrt"])

	//Deletando uma chave
	//Primeirop passar a variável do map e a chave
	delete(user2, "hobs")
	fmt.Println(user2)

	//Inserindo uma chave nova chave ao map
	user2["sign"] = map[string]string{
		"name": "Sagittarius",
	}
	fmt.Println(user2)

}