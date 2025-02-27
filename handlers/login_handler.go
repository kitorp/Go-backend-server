package handlers

import (
	"encoding/json"
	"net"
	"time"

	"../utilities"
)

func loginHandler(conn net.Conn, originalMessage []byte) {

	request := utilities.LoginRequest{}
	err := json.Unmarshal(originalMessage, &request)
	if err != nil {
		Log.WarningF("Json Unmarshal Error. %s", err.Error())
		return
	}

	response := tryLogin(request)
	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. %s", err.Error())
		return
	}

	utilities.Write(conn, dataToSend)
}

func tryLogin(request utilities.LoginRequest) (response utilities.LoginResponse) {

	response.Success = false

	password, id, err := getPasswordAndUserIDFromDB(request.Email)
	if err != nil {
		response.Error = "Database Error"
		return
	}

	passwordMatched, err := ComparePasswords(password, request.Password)
	if err != nil {
		response.Error = "Mismatch Password/Email"
		return
	}
	if passwordMatched {
		token := IssueToken(request.Email, 24*time.Hour)
		err := updateToken(token, request.Email)
		if err != nil {
			response.Error = "Token Error"
			Log.WarningF("Error getting Token. %s", err.Error())
			return
		}
		response.UserID = id
		response.Token = token
		response.Success = true
		return
	} else {
		Log.InfoF("Authentication Error for Email: %s", request.Email)
		response.Error = "Wrong email/password"
		return
	}

}

func authenticateUser(email string, password string, token string, userId int) bool {

	if len(token) >= tokenLength {
		return authenticateUserByToken(token, userId)
	} else if len(email) >= emailLength && len(password) >= passwordLength {
		return authenticateUserByEmailPassword(email, password, userId)
	} else {
		return false
	}

}

func authenticateUserByEmailPassword(email string, password string, userId int) bool {

	req := utilities.LoginRequest{
		Email:    email,
		Password: password,
	}
	response := tryLogin(req)

	if len(response.Error) > 0 {
		return false
	}

	return canUserExecuteRequestOnThisID(email, response.UserID)

}

func authenticateUserByToken(token string, userId int) bool {

	ok, email := verifyTokenAndGetEmail(token)
	if !ok {
		Log.InfoF("Token Verification False for email: %s", email)
		return false
	}
	return canUserExecuteRequestOnThisID(email, userId)

}

func canUserExecuteRequestOnThisID(email string, userId int) bool {

	id, userType, err := getUserIdAndUserTypeFromDB(email)
	if err != nil {
		return false
	}
	if userType == userAdmin || id == userId {
		return true
	}
	return false
}
