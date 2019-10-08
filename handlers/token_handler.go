package handlers

import (
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string
	jwt.StandardClaims
}

func IssueToken(username string, expiraryTime time.Duration) string {

	expirationTime := time.Now().Add(expiraryTime)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		Log.WarningF("Error Issuing Token. ", err.Error())
		return "internal error"
	}

	return tokenString
}

func VerifyToken(tokenString string) (bool, string) {

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

	return true, claims.Username
}
