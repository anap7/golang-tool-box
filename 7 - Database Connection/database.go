package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connString := "user:yourpassword@tcp(localhost:3306)/devbook"
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
