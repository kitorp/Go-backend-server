package handlers

//return error type at line 37

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

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(expiraryTime)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "internal error"
	}

	return tokenString
}

func VerifyToken(token string) (bool, string) {
	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		return false, "Status Unauthorized"
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, "Status Unauthorized"
		}
		return false, "Status Bad Request"
	}

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	return true, claims.Username
}
