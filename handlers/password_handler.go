package handlers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(plainPassword string) (hashedPassword string, err error) {

	pwd := []byte(plainPassword)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		Log.WarningF("Error Hashing Password. ", err.Error())
		return string(hash), err
	}

	return string(hash), nil
}

func ComparePasswords(hashedPassword string, plainPassword string) (correct bool, err error) {

	plainPwd := []byte(plainPassword)
	byteHash := []byte(hashedPassword)
	err = bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		Log.WarningF("Error Comparing Password. ", err.Error())
		return false, err
	}
	return true, nil
}
