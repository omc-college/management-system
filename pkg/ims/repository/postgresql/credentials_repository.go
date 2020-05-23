package postgresql

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/omc-college/management-system/pkg/ims/models"
)

type CredRepository struct {
	db *sqlx.DB
}

func NewCredentialsRepository(dbConnURL string) (*CredRepository, error) {
	db, err := sqlx.Connect("pgx", dbConnURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &CredRepository{
		db: db,
	}, nil
}

//UpdateCredentials
func (cr *CredRepository) UpdateCredentials(c *models.Credentials) error {
	_, err := cr.db.Exec("UPDATE credentials SET password_hash= $1, salt= $2, updated_at=current_timestamp WHERE id= $3", c.PasswordHash, c.Salt, c.ID)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CredRepository) GetCredentialByUserID(usersId string) (*models.Credentials, error) {

	c := &models.Credentials{}
	err := cr.db.Get(&c, "SELECT * FROM credentials WHERE id=$1", usersId)
	if err != nil {
		return nil, err
	}

	return c, nil
}

//InsertCredentials inserts user's credentials into DB
func (cr *CredRepository) InsertCredentials(c *models.Credentials) error {
	_, err := cr.db.Exec("INSERT INTO credentials(id, password_hash, salt, updated_at) VALUES($1, $2, $3, CURRENT_TIMESTAMP)", c.ID, c.PasswordHash, c.Salt)
	if err != nil {
		return err
	}
	return nil
}

func (сr *CredRepository) InsertAccessToken(userId string, accessToken string) error {

	_, err := сr.db.Exec("INSERT INTO user_acsess_tokens VALUES($1, $2, CURRENT_TIMESTAMP)", userId, accessToken)
	if err != nil {
		return err
	}
	return nil
}
