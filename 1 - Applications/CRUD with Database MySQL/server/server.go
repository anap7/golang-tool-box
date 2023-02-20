package server

import (
	"fmt"
	"crud/database"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID uint32 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("Fail to read body request!"))
		return
	}

	var user user
	if error = json.Unmarshal(bodyRequest, &user); error != nil {
		w.Write([]byte("Error to transform user to struct!"))
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error to connect on database!"))
		return
	}
	defer db.Close()
	
	//Prepare Statment
	statement, error := db.Prepare("insert into users (name, email) values (?, ?)")
	if error != nil {
		w.Write([]byte("Error to create the statement!"))
		return
	}
	defer statement.Close()

	insert, error := statement.Exec(user.Name, user.Email)
	if error != nil {
		w.Write([]byte("Error to exec the statement!"))
		return
	}

	idInserted, error := insert.LastInsertId()
	if error != nil {
		w.Write([]byte("Error to get the inserted ID!"))
		return
	}

	//Status Code
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User inserted sucessfully! ID: %d", idInserted)))
}