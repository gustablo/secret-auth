package commons

import "golang.org/x/crypto/bcrypt"

func Compare(hashedPass string, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
}
