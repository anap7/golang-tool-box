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
	router.HandleFunc("/users", server.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", server.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Listening on 5000 Port")
	log.Fatal(http.ListenAndServe(":5000", router))

}