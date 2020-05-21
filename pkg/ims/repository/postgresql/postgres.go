package postgresql

import(
	"fmt"
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/omc-college/management-system/pkg/ims/models"
	"github.com/omc-college/management-system/pkg/db"
)

type SignUpRepository struct {
	DB *sql.DB
}
type CredRepository struct {
	DB *sql.DB
}

func NewSignUpRepository(conf db.RepositoryConfig) (*SignUpRepository, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	db, err := sql.Open("pgx", psqlInfo)

 	return &SignUpRepository{
		DB: db,
	}, err
}
func NewCredentialsRepository(conf db.RepositoryConfig) (*CredRepository, error) {
	pg:= fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	db, err := sql.Open("pgx", pg)

	return &CredRepository{
		DB: db,
	}, err
}

//InsertUser inserts user's data into DB 
func (rep *SignUpRepository)InsertUser(u *models.Users) error {
	_, err := rep.DB.Exec("INSERT INTO users(first_name, last_name, email) VALUES($1, $2, $3)", u.FirstName, u.LastName, u.Email)
	if err != nil {
		return err
	}

	return nil
}

//InsertCredentials inserts user's credentials into DB
func (rep *SignUpRepository)InsertCredentials(c *models.Credentials) error {
	_, err := rep.DB.Exec("INSERT INTO credentials(password_hash, salt, updated_at) VALUES($1, $2, CURRENT_TIMESTAMP)", c.PasswordHash, c.Salt)
	if err != nil {
		return err
	}

	return nil
}

//InsertEmailVerificationToken inserts user's email verification token into DB
func (rep *SignUpRepository)InsertEmailVerificationToken(t *models.EmailVerificationTokens) error {
	_, err := rep.DB.Exec("INSERT INTO email_verification_tokens(verification_token, generated_at) VALUES($1, CURRENT_TIMESTAMP)", t.VerificationToken)
	if err != nil {
		return err
	}
	
	return nil
}

//DeleteEmailVerificationToken deletes user's token after verification
func (rep *SignUpRepository)DeleteEmailVerificationToken(t *models.EmailVerificationTokens) error {
	_, err := rep.DB.Exec("DELETE FROM email_verification_tokens WHERE verification_token = $1", t.VerificationToken)
    if err != nil {
        return err
	}
	
	return nil
}

//UpdateUser updates user's verified status 
func (rep *SignUpRepository)SetUserVerified(u *models.Users) error {
	_, err := rep.DB.Exec("UPDATE users SET verified = true WHERE email = $1", u.Email)
    if err != nil {
		return err
	}
	
	return nil
}

//UpdateCredentials
func (repository *CredRepository) UpdateCredentials(c *models.Credentials) error {
	_, err := repository.DB.Exec("UPDATE credentials SET password_hash= $1,salt=$2,updated_at=current_timestamp WHERE id= $3", c.PasswordHash, c.Salt)
	if err != nil {
		return err
	}

	return nil
}
