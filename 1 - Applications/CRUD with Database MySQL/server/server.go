package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID uint32 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

//Create a new user on database, expected name and email
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("Fail to read body request!"))
		return
	}

	var user user
	if error = json.Unmarshal(bodyRequest, &user); error != nil {
		w.Write([]byte("Error to transform user to struct!"))
		return
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

//Get all users from database
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error to connect on database!"))
		return
	}
	defer db.Close()

	//Prepare query
	rows, error := db.Query("select * from users")
	if error != nil {
		w.Write([]byte("Error to search users from database!"))
		return
	}
	defer rows.Close()

	//Slice de usuários
	var users []user

	//Lendo um por um
	for rows.Next() {
		var user user

		if error := rows.Scan(&user.ID, &user.Name, &user.Email); error != nil {
			w.Write([]byte("Error to scan user"))
			return
		}

		users = append(users, user)
	}

	//Status Code
	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(users); error != nil {
		w.Write([]byte("Error to transform slice users in JSON"))
		return 
	}
}

//Get an user from database, expected ID for search
func GetUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	//Nome do parametro, base numérica (decimal) e tipo int32 que foi declarado no struct
	ID, error := strconv.ParseUint(param["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Error to transform the request param to integer"))
		return 
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error to connect on database!"))
		return
	}
	defer db.Close()

	//Prepare query
	row, error := db.Query("select * from users where id = ?", ID)
	if error != nil {
		w.Write([]byte("Error to search the user from database!"))
		return
	}

	var user user
	if row.Next() {
		if error := row.Scan(&user.ID, &user.Name, &user.Email); error != nil {
			w.Write([]byte("Error to scan user"))
			return
		}
	}

	//Status Code
	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(user); error != nil {
		w.Write([]byte("Error to transform user in JSON"))
		return 
	}
}

//Update user data on database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	//Nome do parametro, base numérica (decimal) e tipo int32 que foi declarado no struct
	ID, error := strconv.ParseUint(param["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Error to transform the request param to integer"))
		return 
	}

	bodyRequest, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("Fail to read body request!"))
		return
	}

	var user user
	if error = json.Unmarshal(bodyRequest, &user); error != nil {
		w.Write([]byte("Error to transform user to struct!"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error to connect on database!"))
		return
	}
	defer db.Close()

	//Prepare Statment
	statement, error := db.Prepare("update users set name = ?, email = ? where id = ?")
	if error != nil {
		w.Write([]byte("Error to create the statement!"))
		return
	}
	defer statement.Close()

	if _, error := statement.Exec(user.Name, user.Email, ID); error != nil {
		w.Write([]byte("Error to update the user!"))
		return
	}

	//Status Code
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("User updated sucessfully! ID: %d", ID)))
}

//Delete user data from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	//Nome do parametro, base numérica (decimal) e tipo int32 que foi declarado no struct
	ID, error := strconv.ParseUint(param["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Error to transform the request param to integer"))
		return 
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error to connect on database!"))
		return
	}
	defer db.Close()

	//Prepare Statment
	statement, error := db.Prepare("delete from users where id = ?")
	if error != nil {
		w.Write([]byte("Error to create the statement!"))
		return
	}
	defer statement.Close()

	if _, error := statement.Exec(ID); error != nil {
		w.Write([]byte("Error to delete the user!"))
		return
	}

	//Status Code
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("User deleted sucessfully! ID: %d", ID)))
}


