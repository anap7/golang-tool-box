package main

import (
	"fmt"
	"module/helper"
	//Importando o arquivo externo
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main file")
	helper.WriteFromFilePointUppercase()
	//Esperando um erro de retorno, pois estou enviando um formato diferente de e-mail
	error := checkmail.ValidateFormat("123")
	fmt.Println(error);
}