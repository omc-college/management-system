package postgresql

import (
	"database/sql"
	"errors"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/omc-college/management-system/pkg/ims/models"
)

//
type UsersRepository struct {
	Db *sql.DB
}

// Create new user repository
func NewUsersRepository(dbPath string) (*UsersRepository, error) {
	Db, err := sql.Open("pgx", dbPath)
	if err != nil {
		return nil, err
	}
	err = Db.Ping()
	if err != nil {
		return nil, err
	}
	return &UsersRepository{Db: Db}, nil
}

func GetAllUsers(repository *UsersRepository) ([]models.Users, error) {
	query := `SELECT *  FROM users;`

	var users []models.Users

	rows, err := repository.Db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var u = models.Users{}
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.MobilePhone, &u.Role, &u.CreatedAt, &u.ModifiedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(repository *UsersRepository, ID int) (models.Users, error) {
	query := `SELECT * FROM users WHERE id = $1;`

	var user = models.Users{}

	rows, err := repository.Db.Query(query, ID)
	if err != nil {
		return models.Users{}, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.MobilePhone, &user.Role, &user.CreatedAt, &user.ModifiedAt)
		if err != nil {
			return models.Users{}, err
		}
	}

	if user.ID == 0 {
		return models.Users{}, ErrNoRows
	}

	err = rows.Err()
	if err != nil {
		return models.Users{}, err
	}
	return user, nil
}

func AddUser(repository *UsersRepository, user models.Users) error {
	query := `INSERT INTO users (first_name, last_name, email, mobile_phone, role, created_at, modified_at) 
    VALUES ($1, $2, $3, $4, $5, current_timestamp, null ) RETURNING (id)`

	var uID int

	err := repository.Db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.MobilePhone, user.Role).Scan(&uID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(repository *UsersRepository, user models.Users) error {
	query := `SELECT FROM users WHERE id = $1`

	err := repository.Db.QueryRow(query, user.ID).Scan()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrNoRows
		} else {
			err = QueryError{queryErrorMessage, err}
		}
		return err
	}

	query = `UPDATE users SET first_name = $1, last_name= $2, email= $3, mobile_phone= $4, modified_at= current_timestamp WHERE id = $5`

	_, err = repository.Db.Exec(query, user.FirstName, user.LastName, user.Email, user.MobilePhone, user.Role, user.ID)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	return nil
}

func DeleteUser(repository *UsersRepository, UserID int) error {
	query := `SELECT FROM users WHERE id = $1`

	err := repository.Db.QueryRow(query, UserID).Scan()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrNoRows
		} else {
			err = QueryError{queryErrorMessage, err}
		}
		return err
	}

	query = `DELETE FROM users WHERE id = $1`

	_, err = repository.Db.Exec(query, UserID)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	return nil
}
