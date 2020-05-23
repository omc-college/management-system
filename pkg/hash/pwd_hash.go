package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func clear(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0;
	}
}

//PwdHash returns password hash which generates with bcrypt
func PwdHash(password string) (string, error) {
	pwd := []byte(password)
	defer clear(pwd)
	hash, err := bcrypt.GenerateFromPassword(pwd, 15)
	return string(hash), err
}