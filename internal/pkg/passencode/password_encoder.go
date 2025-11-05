package passencode

import (
	"golang.org/x/crypto/bcrypt"
)

const DefaultCost = 11

func HashPassword(password string) (bool, string) {

	var hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	if err != nil {
		return false, password
	}

	return true, string(hashedPassword)
}

func ComparePasswords(hashedPwd string, plainPwd string) bool {

	byteHash := []byte(hashedPwd)
	bytePassword := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		return false
	}
	return true
}
