package postgresql

import (
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/pgtype"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/omc-college/management-system/pkg/ims/models"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(dbConnURL string) (*UsersRepository, error) {
	db, err := sqlx.Connect("pgx", dbConnURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &UsersRepository{
		db: db,
	}, nil
}

func (ur *UsersRepository) GetAllUsers() ([]models.Users, error) {

	var users []models.Users

	err := ur.db.Select(&users, "SELECT * FROM users ORDER BY id ASC")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UsersRepository) GetUser(ID int) (*models.Users, error) {

	u := user{}

	err := ur.db.Get(&u, "SELECT * FROM users WHERE id= $1", ID)
	if err != nil {
		return nil, err
	}

	if u.ID == 0 {
		return nil, ErrNoRows
	}

	return ToUser(u)
}

func (ur *UsersRepository) UpdateUser(u models.Users) error {
	user := user{}
	err := ur.db.Get(&user, "SELECT * FROM users WHERE id= $1", u.ID)
	if err != nil {
		return err
	}

	_, err = ur.db.Exec("UPDATE users SET first_name = $1, last_name= $2, email= $3, mobile_phone= $4, modified_at= CURRENT_TIMESTAMP, roles= $5 WHERE id = $6", u.FirstName, u.LastName, u.Email, u.MobilePhone, u.Roles, u.ID)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	return nil
}

func (ur *UsersRepository) DeleteUser(UserID int) error {

	_, err := ur.db.Exec("DELETE FROM users WHERE id= $1", UserID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UsersRepository) GetUserByEmail(email string) (*models.Users, error) {

	u := user{}

	err := ur.db.Get(&u, "SELECT * FROM users WHERE email= $1", email)
	if err != nil {
		return nil, err
	}

	return ToUser(u)
}

func ToUser(privatUser user) (*models.Users, error) {

	roles := []string{}
	err := privatUser.Roles.AssignTo(&roles)
	if err != nil {
		return nil, err
	}
	var genericUser *models.Users
	genericUser = &models.Users{
		ID:          privatUser.ID,
		FirstName:   privatUser.FirstName,
		LastName:    privatUser.LastName,
		Email:       privatUser.Email,
		MobilePhone: privatUser.MobilePhone,
		CreatedAt:   privatUser.CreatedAt,
		ModifiedAt:  privatUser.ModifiedAt,
		Roles:       roles,
		Verified:    privatUser.Verified,
	}
	return genericUser, nil
}
