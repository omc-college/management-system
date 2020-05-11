package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	 "github.com/omc-college/management-system/book/route"
	"io/ioutil"
	"net/http"
)

//Book Struct (Model)
type Book struct {
	ID    string `json:"ID"`
	Title string `json:"Title"`
}

//Init database and error vars
var db *sql.DB
var err error
var books []Book

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []Book

	result, err := db.Query("SELECT id,title  from Books")
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
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params

	result, err := db.Query("SELECT  FROM Books WHERE ID = $1", params["ID"])
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
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	result, err := db.Query("INSERT INTO Books(ID,Title) VALUES(2,2)")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	fmt.Fprintf(w, "New book was created")

}

//Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	result, err := db.Query("UPDATE Books SET Title = $3  WHERE ID = $1")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	fmt.Fprintf(w, "Book with ID = %s was updated", params["ID"])

}

//Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := db.Query("DELETE FROM Books WHERE ID = $1")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	fmt.Fprintf(w, "Book with ID = %s was deleted", params["ID"])

}

