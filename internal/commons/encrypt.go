package commons

import "golang.org/x/crypto/bcrypt"

func Encrypt(str *string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(*str), 14)

	*str = string(bytes)
}
