package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Sub   string   `json:"sub"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
	jwt.StandardClaims
}
