package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//connectionString := "golang:Golang.17386/devbook?charset=utf8&parseTime=True&loc=Local"
	//connString := "golanguser:Golang17396@@tcp@(127.0.0.1:3306)/devbook"
	connString := "golang:Golang_17396@tcp(localhost:3306)/devbook"
	db, error := sql.Open("mysql", connString)
	if error != nil {
		log.Fatal(error)
	}
	defer db.Close()

	if error = db.Ping(); error != nil {
		fmt.Println("Caiu no erro!")
		log.Fatal(error)
	}

	fmt.Println("Conexão show!")

	rows, error := db.Query("select * from usuarios")
	if error != nil {
		log.Fatal(error)
	}
	defer rows.Close()

	fmt.Println(rows)
}
