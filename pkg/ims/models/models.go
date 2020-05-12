package models

import "github.com/dgrijalva/jwt-go"

type Users struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Verified  bool   `json:"verified"`
}

type Claims struct {
	Email string `json:"email"`
	Sub   string `json:"sub"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type Credentials struct {
	PasswordHash string
	Salt string
}

type EmailVerificationTokens struct {
	VerificationToken string
}