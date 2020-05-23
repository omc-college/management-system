package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Users struct {
	ID          int      `json:"id"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Email       string   `json:"email"`
	MobilePhone string   `json:"mobile_phone"`
	CreatedAt   string   `json:"created_at"`
	ModifiedAt  string   `json:"modified_at"`
	Roles       []string `json:"roles"`
	Verified    bool     `json:"verified"`
}

type Claims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type Credentials struct {
	ID           string `json:"id"`
	PasswordHash string `json:"password_hash"`
	Salt         string `json:"salt"`
	UpdatedAt    string `json:"updated_at"`
}

type EmailVerificationTokens struct {
	ID                string `json:"id"`
	VerificationToken string `json:"verification_token"`
	GeneretedAt       string `json:"generated_at"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
