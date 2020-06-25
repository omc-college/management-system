package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID          int       `json:"id" db:"id"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	Email       string    `json:"email" db:"email"`
	MobilePhone string    `json:"mobile_phone" db:"mobile_phone"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	ModifiedAt  time.Time `json:"modified_at" db:"modified_at"`
	Roles       []string  `json:"roles" db:"roles"`
	Verified    bool      `json:"verified" db:"verified"`
}

type Claims struct {
	Sub   string   `json:"sub"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

type Credentials struct {
	ID           int       `json:"id" db:"id"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	Salt         string    `json:"salt" db:"salt"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type EmailVerificationTokens struct {
	ID                int    `json:"id" db:"id"`
	VerificationToken string `json:"verification_token" db:"verification_token"`
	GeneretedAt       string `json:"generated_at" db:"generated_at"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SignupRequest struct {
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
type ResetRequest struct{
	Email   string `json:"email" db:"email"`
}
type CheckResetToken struct{
	Token EmailVerificationTokens `json:"token" db:"token"`
}
type ConfirmResetRequest struct{
	ID       int 	 `json:"id" db:"id"`
	Password string `json:"password" db:"password"`
}
