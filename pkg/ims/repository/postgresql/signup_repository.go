package postgresql

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/omc-college/management-system/pkg/ims/models"
)

type SignUpRepository struct {
	db *sqlx.DB
}

func NewSignUpRepository(dbConnURL string) (*SignUpRepository, error) {
	db, err := sqlx.Connect("pgx", dbConnURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &SignUpRepository{
		db: db,
	}, nil
}

//InsertUser inserts user's data into DB
func (sr *SignUpRepository) InsertUser(u *models.Users) error {
	_, err := sr.db.Exec("INSERT INTO users(id,first_name, last_name, email, mobile_phone, created_at, modified_at, roles) VALUES($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $6)", u.ID, u.FirstName, u.LastName, u.Email, u.MobilePhone, u.Roles)
	if err != nil {
		return err
	}

	return nil
}

//InsertEmailVerificationToken inserts user's email verification token into DB
func (sr *SignUpRepository) InsertEmailVerificationToken(t *models.EmailVerificationTokens) error {
	_, err := sr.db.Exec("INSERT INTO email_verification_tokens(id,verification_token, generated_at) VALUES($1, $2, CURRENT_TIMESTAMP)", t.ID, t.VerificationToken)
	if err != nil {
		return err
	}

	return nil
}

//DeleteEmailVerificationToken deletes user's token after verification
func (sr *SignUpRepository) DeleteEmailVerificationToken(t *models.EmailVerificationTokens) error {
	_, err := sr.db.Exec("DELETE FROM email_verification_tokens WHERE verification_token = $1", t.VerificationToken)
	if err != nil {
		return err
	}

	return nil
}

//UpdateUser updates user's verified status
func (sr *SignUpRepository) SetUserVerified(u *models.Users) error {
	_, err := sr.db.Exec("UPDATE users SET verified = true WHERE email = $1", u.Email)
	if err != nil {
		return err
	}

	return nil
}
