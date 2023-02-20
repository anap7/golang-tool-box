package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver connection with MySQL
)

//Open Connection with database (*A função retorna dois valores: o ponteiro de conexão do sql e um erro)
func Connect() (*sql.DB, error) {
	connectionString := "golang:Golang_17386/devbook?charset=utf8&parseTime=True&loc=Local" 
	db, error := sql.Open("mysql", connectionString)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		fmt.Println("Problem to ping database!")
		return nil, error
	}

	fmt.Println("Connection is Open succesfully!")

	return db, nil
}