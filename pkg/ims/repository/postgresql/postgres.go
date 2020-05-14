package postgresql

import (
	"database/sql"

	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"

)

//
////Init database and error vars
//type PasswordRepository struct {
//	db*sql.DB
//}



var Db *sql.DB
var err error

func UpdatePassword(id int){

	result, err := Db.Query("UPDATE credentials SET password_hash = $2,salt = $3, updated_at=current_timestamp WHERE id = $1 returning id,password_hash,salt,updated_at", id )
	if err != nil {
		return QueryError {UpdateCredentialsErrorMessage, err}
	}
	defer result.Close()
}