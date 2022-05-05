package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Bcrypt(pwd string) string {
	password := []byte(pwd)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func IsBcryptMatched(pwd string, hashedPWD string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPWD), []byte(pwd))
	if err == nil {
		return true
	} else {
		return false
	}
}
