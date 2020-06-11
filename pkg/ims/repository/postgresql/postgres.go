package dbquery

import (
	"database/sql"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"log"
)


func DbQuerier(tableToQuery string) (*sql.Rows, error) {
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	var dbqrows *sql.Rows
	dbqrows, err = db.Query("SELECT * FROM "+ tableToQuery) //query из бази в rows
	if err != nil {
		log.Fatal(err)
	}
	defer dbqrows.Close()
	return dbqrows,err
}