package helper

import "golang.org/x/crypto/bcrypt"

func BcryptPassword(password string) string {
	passwordBytes := []byte(password)
	hashedPassBytes, err := bcrypt.GenerateFromPassword(passwordBytes, 12)
	DoPanicIfError(err)
	hashedPassStr := string(hashedPassBytes)
	return hashedPassStr
}
