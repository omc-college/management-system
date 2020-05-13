package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/omc-college/management-system/book/postgres"
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
	r.HandleFunc("/books", postgres.GetBooks).Methods("GET")
	r.HandleFunc("/books/{ID}", postgres.GetBook).Methods("GET")
	r.HandleFunc("/books", postgres.CreateBook).Methods("POST")
	r.HandleFunc("/books/{ID}", postgres.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{ID}", postgres.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8001", r))
}
