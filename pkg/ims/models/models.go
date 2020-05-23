package models

import "github.com/dgrijalva/jwt-go"

type Users struct {
	ID          int    `json:"ID"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	MobilePhone string `json:"mobile_phone"`
	CreatedAt   string `json:"created_at"`
	ModifiedAt  string `json:"modified_at"`
	Role        string `json:"role"`
	Verified    bool   `json:"verified"`
}

type Claims struct {
	Email string `json:"email"`
	Sub   string `json:"sub"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type Credentials struct {
	PasswordHash string
	Salt         string
}

type EmailVerificationTokens struct {
	VerificationToken string `json:"verification_token"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
