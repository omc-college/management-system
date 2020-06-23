package postgresql

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/omc-college/management-system/pkg/ims/models"
)

type CredRepository struct {
	db sqlx.Ext
}

func NewCredentialsRepository(db sqlx.Ext) *CredRepository {
	return &CredRepository{
		db: db,
	}
}

//UpdateCredentials
func (cr *CredRepository) UpdateCredentials(c *models.Credentials, userId int) error {
	_, err := cr.db.Exec("UPDATE credentials SET password_hash= $1, salt= $2, updated_at=current_timestamp WHERE id= $3", c.PasswordHash, c.Salt, userId)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CredRepository) GetCredentialByUserID(usersId string) (*models.Credentials, error) {
	c := &models.Credentials{}

	result, err := cr.db.Query("SELECT * FROM credentials WHERE id=$1", usersId)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		err = result.Scan(&c.ID, &c.PasswordHash, &c.Salt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

//InsertCredentials inserts user's credentials into DB
func (cr *CredRepository) InsertCredentials(c *models.Credentials) error {
	_, err := cr.db.Exec("INSERT INTO credentials(id, password_hash, salt, updated_at) VALUES($1, $2, $3, CURRENT_TIMESTAMP)", c.ID, c.PasswordHash, c.Salt)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	return nil
}

func (сr *CredRepository) InsertAccessToken(userId string, accessToken string) error {

	_, err := сr.db.Exec("INSERT INTO user_access_tokens VALUES($1, $2, CURRENT_TIMESTAMP)", userId, accessToken)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CredRepository) UpdateAccessToken(userId string, accessToken string) error {
	_, err := cr.db.Exec("UPDATE user_access_tokens SET  access_token= $1, generated_at=current_timestamp WHERE id= $2", accessToken, userId)
	if err != nil {
		return err
	}

	return nil
}
