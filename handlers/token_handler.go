package handlers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Email string
	jwt.StandardClaims
}

func IssueToken(username string, expiraryTime time.Duration) string {

	expirationTime := time.Now().Add(expiraryTime)

	claims := &Claims{
		Email: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		Log.WarningF("Error Issuing Token. %s", err.Error())
		return "internal error"
	}

	return tokenString
}

func verifyTokenAndGetEmail(tokenString string) (IsCorrectToken bool, email string) {

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !token.Valid {
		Log.WarningF("Token Not Valid.")
		return false, "Status Unauthorized"
	}
	if err != nil {
		Log.WarningF("Token Error.", err.Error())
		if err == jwt.ErrSignatureInvalid {
			return false, "Status Unauthorized"
		}
		return false, "Status Bad Request"
	}

	return true, claims.Email
}
