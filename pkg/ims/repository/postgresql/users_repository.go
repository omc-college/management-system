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

func GetAllUsers(repository *UsersRepository) ([]models.User, error) {
	query := `SELECT *  FROM users;`

	var users []models.User

	rows, err := repository.Db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var u = models.User{}
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.MobilePhone, &u.CreatedAt, &u.ModifiedAt)
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

func GetUser(repository *UsersRepository, UserID int) (models.User, error) {
	query := `SELECT * FROM users WHERE userid = $1;`

	var user = models.User{}

	rows, err := repository.Db.Query(query, UserID)
	if err != nil {
		return models.User{}, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.MobilePhone, &user.CreatedAt, &user.ModifiedAt)
		if err != nil {
			return models.User{}, err
		}
	}

	if user.ID == 0 {
		return models.User{}, ErrNoRows
	}

	err = rows.Err()
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func AddUser(repository *UsersRepository, user models.User) error {
	query := `INSERT INTO users (firstname, lastname, email, mobilephone, createdAt, modifiedAt) 
    VALUES ($1, $2, $3, $4, current_timestamp, current_timestamp) RETURNING (userid)`

	var uID int

	err := repository.Db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.MobilePhone).Scan(&uID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(repository *UsersRepository, user models.User, UserID int) error {
	query := `SELECT FROM users WHERE userid = $1`

	err := repository.Db.QueryRow(query, UserID).Scan()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrNoRows
		} else {
			err = QueryError{queryErrorMessage, err}
		}
		return err
	}

	query = `UPDATE users SET firstname = $1, lastname= $2, email= $3, mobilephone= $4, modifiedat= current_timestamp WHERE userid = $5`

	_, err = repository.Db.Exec(query, user.FirstName, user.LastName, user.Email, user.MobilePhone, UserID)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	return nil
}

func DeleteUser(repository *UsersRepository, UserID int) error {
	query := `SELECT FROM users WHERE userid = $1`

	err := repository.Db.QueryRow(query, UserID).Scan()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrNoRows
		} else {
			err = QueryError{queryErrorMessage, err}
		}
		return err
	}

	query = `DELETE FROM users WHERE userid = $1`

	_, err = repository.Db.Exec(query, UserID)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	return nil
}
