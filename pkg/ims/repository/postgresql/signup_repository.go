package postgresql

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/omc-college/management-system/pkg/ims/models"
)

type SignUpRepository struct {
	db sqlx.Ext
}

func NewSignUpRepository(db sqlx.Ext) *SignUpRepository {
	return &SignUpRepository{
		db: db,
	}
}

//InsertEmailVerificationToken inserts user's email verification token into DB
func (sr *SignUpRepository) InsertEmailVerificationToken(t *models.EmailVerificationTokens) error {
	_, err := sr.db.Exec("INSERT INTO email_verification_tokens(id,verification_token, generated_at) VALUES($1, $2, CURRENT_TIMESTAMP)", t.ID, t.VerificationToken)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	return nil
}

//DeleteEmailVerificationToken deletes user's token after verification
func (sr *SignUpRepository) DeleteEmailVerificationToken(t *models.EmailVerificationTokens) error {
	_, err := sr.db.Exec("DELETE FROM email_verification_tokens WHERE verification_token = $1", t.VerificationToken)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	return nil
}

//UpdateUser updates user's verified status
func (sr *SignUpRepository) SetUserVerified(u *models.User) error {
	_, err := sr.db.Exec("UPDATE users SET verified = true WHERE email = $1", u.Email)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	return nil
}
