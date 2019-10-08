package handlers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(plainPassword string) (hashedPassword string, err error) {

	passwordInByte := []byte(plainPassword)
	hash, err := bcrypt.GenerateFromPassword(passwordInByte, bcrypt.MinCost)
	if err != nil {
		Log.WarningF("Error Hashing Password. %s", err.Error())
		return string(hash), err
	}
	return string(hash), nil
}

func ComparePasswords(hashedPassword string, plainPassword string) (correct bool, err error) {

	hashedPasswordInByte := []byte(hashedPassword)
	plainPasswordInByte := []byte(plainPassword)

	err = bcrypt.CompareHashAndPassword(hashedPasswordInByte, plainPasswordInByte)
	if err != nil {
		Log.WarningF("False Password. %s", err.Error())
		return false, err
	}
	return true, nil
}
