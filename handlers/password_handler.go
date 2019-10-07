package handlers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwdString string) string {

	pwd := []byte(pwdString)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func ComparePasswords(hashedPwd string, plainpass string) bool {

	plainPwd := []byte(plainpass)
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}
