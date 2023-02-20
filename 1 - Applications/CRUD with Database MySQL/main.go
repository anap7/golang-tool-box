package main

import (
	"fmt"
	"log"
	"net/http"
	"crud/server"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	/*Chamando a função de criar usuário que está no pacote server*/
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)

	fmt.Println("Listening on 5000 Port")
	log.Fatal(http.ListenAndServe(":5000", router))

}