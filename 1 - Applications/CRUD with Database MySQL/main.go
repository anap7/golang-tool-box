package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fmt.Println("Listening on 5000 Port")
	log.Fatal(http.ListenAndServe(":5000", router))

}