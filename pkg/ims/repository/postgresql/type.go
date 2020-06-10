package postgresql

import (
	"github.com/jackc/pgx/pgtype"
	"time"
)

type user struct {
	ID          int              `json:"id" db:"id"`
	FirstName   string           `json:"first_name" db:"first_name"`
	LastName    string           `json:"last_name" db:"last_name"`
	Email       string           `json:"email" db:"email"`
	MobilePhone string           `json:"mobile_phone" db:"mobile_phone"`
	CreatedAt   time.Time        `json:"created_at" db:"created_at"`
	ModifiedAt  time.Time        `json:"modified_at" db:"modified_at"`
	Roles       pgtype.TextArray `json:"roles" db:"roles"`
	Verified    bool             `json:"verified" db:"verified"`
}
