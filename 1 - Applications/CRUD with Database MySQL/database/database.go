package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver connection with MySQL
)

//Open Connection with database (*A função retorna dois valores: o ponteiro de conexão do sql e um erro)
func Connect() (*sql.DB, error) {
	connectionString := "golang:Golang_17396@tcp(localhost:3306)/devbook"
	db, error := sql.Open("mysql", connectionString)
	if error != nil {
		fmt.Println("Problem to open database!")
		fmt.Print(error)
		return nil, error
	}

	if error = db.Ping(); error != nil {
		fmt.Println("Problem to ping database!")
		fmt.Print(error)
		return nil, error
	}
	
	return db, nil
}