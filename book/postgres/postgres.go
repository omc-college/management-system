package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

//Book Struct (Model)
type Book struct {
	ID    string `json:"ID"`
	Title string `json:"Title"`
}

//Init database and error vars
var Db *sql.DB
var Err error
var books []Book

//Get All Books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []Book

	result, err := Db.Query("SELECT id,title  from Books")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()
	for result.Next() {

		var book Book

		err := result.Scan(&book.ID, &book.Title)
		if err != nil {
			panic(err.Error())
		}

		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

//Get Single Book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params

	result, err := Db.Query("SELECT id,title FROM Books WHERE ID = $1", params["ID"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var book Book

	for result.Next() {
		err := result.Scan(&book.ID, &book.Title)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(book)
}

//Create a New Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	result, err := Db.Query("INSERT INTO Books(ID,Title) VALUES($1,$1)")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	fmt.Fprintf(w, "New book was created")

}

//Update Book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	result, err := Db.Query("UPDATE Books SET Title = $1  WHERE ID = $1")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	fmt.Fprintf(w, "Book with ID = %s was updated", params["ID"])

}

//Delete Book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := Db.Query("DELETE FROM Books WHERE ID = $1")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	fmt.Fprintf(w, "Book with ID = %s was deleted", params["ID"])

}
