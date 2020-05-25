package pwd

import (
	"golang.org/x/crypto/bcrypt"
)

func clear(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0;
	}
}

// Hash returns password hash which generates with bcrypt
func Hash(password string, salt string) (string, error) {
	pwd := []byte(password + salt)
	defer clear(pwd)
	hash, err := bcrypt.GenerateFromPassword(pwd, 15)
	return string(hash), err
}