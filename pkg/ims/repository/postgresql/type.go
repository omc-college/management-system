package postgresql

import (
	"github.com/jackc/pgx/pgtype"
)

type user struct {
	ID          int              `json:"id"`
	FirstName   string           `json:"first_name"`
	LastName    string           `json:"last_name"`
	Email       string           `json:"email"`
	MobilePhone string           `json:"mobile_phone"`
	CreatedAt   string           `json:"created_at"`
	ModifiedAt  string           `json:"modified_at"`
	Roles       pgtype.TextArray `json:"role"`
	Verified    bool             `json:"verified"`
}
