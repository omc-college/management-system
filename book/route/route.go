package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/omc-college/management-system/book/postgres"
	"log"
	"net/http"
)

func main() {

	db, err = sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	//Init Router

	r = mux.NewRouter()

	//Route Handlers/Endpoints
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{ID}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{ID}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{ID}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8001", r))
}