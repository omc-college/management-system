package postgresql

import (
	_ "github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	_ "github.com/jackc/pgx/pgtype"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/omc-college/management-system/pkg/ims/models"
)

type UsersRepository struct {
	db sqlx.Ext
}

func NewUsersRepository(db sqlx.Ext) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

//InsertUser inserts user's data into DB
func (ur *UsersRepository) InsertUser(u *models.User) error {
	var err error

	roles := pgtype.TextArray{}
	err = roles.Set(u.Roles)
	if err != nil {
		return err
	}

	result, err := ur.db.Queryx("INSERT INTO users (first_name, last_name, email, mobile_phone, created_at, modified_at, roles) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $5) RETURNING id", u.FirstName, u.LastName, u.Email, u.MobilePhone, roles.Get())
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	for result.Next() {
		err = result.Scan(&u.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ur *UsersRepository) GetAllUsers() ([]models.User, error) {

	var users []models.User

	result, err := ur.db.Queryx("SELECT * FROM users ORDER BY id ASC")
	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var u user

		err := result.StructScan(&u)
		if err != nil {
			return users, err
		}

		us, err := ToUser(u)
		if err != nil {
			return users, err
		}

		users = append(users, *us)
	}

	return users, nil
}

func (ur *UsersRepository) GetUser(ID int) (*models.User, error) {

	u := user{}

	result, err := ur.db.Queryx("SELECT * FROM users WHERE id= $1", ID)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		err = result.StructScan(&u)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, ErrNoRows
	}

	return ToUser(u)
}

func (ur *UsersRepository) UpdateUser(u models.User) error {

	user := user{}

	result, err := ur.db.Queryx("SELECT * FROM users WHERE id= $1", u.ID)
	if err != nil {
		return err
	}

	for result.Next() {
		err = result.StructScan(&user)
		if err != nil {
			return err
		}
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

func (ur *UsersRepository) GetUserByEmail(email string) (*models.User, error) {
	u := user{}

	result, err := ur.db.Queryx("SELECT * FROM users WHERE email= $1", email)
	if err != nil {
		return nil, QueryError{queryErrorMessage, err}
	}

	for result.Next() {
		err = result.StructScan(&u)
		if err != nil {
			return nil, err
		}
	}

	return ToUser(u)
}

// GetEmailByToken returns user's email by token
func (ur *UsersRepository) GetEmailByToken (token *models.EmailVerificationTokens) (string, error) {
	var email string

	result, err := ur.db.Queryx("SELECT (SELECT email FROM users WHERE id = email_verification_tokens.id) FROM email_verification_tokens WHERE verification_token = $1", token.VerificationToken)
	if err != nil {
		return email, QueryError{queryErrorMessage, err}
	}

	for result.Next() {
		err = result.Scan(&email)
		if err != nil {
			return email, err
		}
	}

	return email, nil
}

func ToUser(u user) (*models.User, error) {

	roles := []string{}

	if u.Roles.Status != pgtype.Undefined {
		err := u.Roles.AssignTo(&roles)
		if err != nil {
			return nil, err
		}
	}

	var genericUser *models.User

	genericUser = &models.User{
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		MobilePhone: u.MobilePhone,
		CreatedAt:   u.CreatedAt,
		ModifiedAt:  u.ModifiedAt,
		Roles:       roles,
		Verified:    u.Verified,
	}

	return genericUser, nil
}
