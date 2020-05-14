package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/omc-college/management-system/password/postgres"
	"log"
	"net/http"
)

func main() {
	postgres.Db, postgres.Err = sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if postgres.Err != nil {
		panic(postgres.Err.Error())
	}

	defer postgres.Db.Close()

	//Init Router
	r := mux.NewRouter()

	//Route Handlers/Endpoints
	r.HandleFunc("/password", postgres.Getpassword).Methods("GET")
	r.HandleFunc("/password", postgres.Insertpassword).Methods("POST")
	r.HandleFunc("/password", postgres.DeleteLastPassword).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8001", r))
}
