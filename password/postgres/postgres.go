package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
)

//credentials Struct (Model)
type credentials struct {
	passwordHash    string
	salt            string
	currentPassword string
	newPassword     string
	token           string
	newtoken        string
}

//Init database and error vars
var Db *sql.DB
var Err error
var password []credentials

//Get All Books
func Getpassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var password []credentials

	result, err := Db.Query("SELECT current_password FROM  password WHERE id=4")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()
	for result.Next() {

		var password credentials

		err := result.Scan(&password.currentPassword)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(&password)
}

//Insert a New password
func Insertpassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	result, err := Db.Query("INSERT INTO password (id, current_password, new_password, token, new_token, verified, salt, password_hash) VALUES ('1','decer','12eedw','21212','decer','off','323dd32s23','fdcsdfscxd')")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	fmt.Fprintf(w, "New password was created")

}

//Delete Password after change
func DeleteLastPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := Db.Query("DELETE  FROM password  WHERE ID = 1")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	fmt.Fprintf(w, "Password with ID = %s was deleted", params["ID"])

}
