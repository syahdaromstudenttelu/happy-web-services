package helper

import "golang.org/x/crypto/bcrypt"

func ValidatePassword(hashedPassword, password string) error {
	passwordBytes := []byte(password)
	hashedPassBytes := []byte(hashedPassword)
	return bcrypt.CompareHashAndPassword(hashedPassBytes, passwordBytes)
}
