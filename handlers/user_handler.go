package handlers

import (
	"encoding/json"
	"net"

	"../utilities"

	"github.com/badoux/checkmail"
)

func createUserHandler(conn net.Conn, data []byte) {

	request := utilities.CreateUserRequest{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		Log.WarningF("Json Unmarshal Error. %s", err.Error())
		return
	}

	response := utilities.CommonResponse{
		Success: false,
	}

	if authenticateUser(request.Email, request.Password, request.Token, 0) {
		if !isEmailValid(request.Email) {
			response.Error = "Email not valid"
			return
		}
		err := createUser(request.UserEmail, request.UserPassword, request.UserType)
		if err != nil {
			response.Error = "Error creating user"
		} else {
			response.Success = true
		}

	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. %s", err.Error())
		return
	}

	utilities.Write(conn, dataToSend)
}

func listUserHandler(conn net.Conn, data []byte) {

	request := utilities.ListUserRequest{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		Log.WarningF("Json Unmarshal Error. %s", err.Error())
		return
	}

	response := utilities.ListUserResponse{
		Success: false,
	}

	if authenticateUser(request.Email, request.Password, request.Token, 0) {
		list, err := listUser(request.Limit, request.Offset)
		if err != nil {
			response.Error = "Error listing user"
		} else {
			response.Users = append(response.Users, list...)
			response.Success = true
		}

	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. %s", err.Error())
		return
	}

	utilities.Write(conn, dataToSend)
}

func deleteUserHandler(conn net.Conn, data []byte) {

	request := utilities.DeleteUserRequest{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		Log.WarningF("Json Unmarshal Error. %s", err.Error())
		return
	}

	response := utilities.CommonResponse{
		Success: false,
	}
	if authenticateUser(request.Email, request.Password, request.Token, 0) {
		err := deleteUser(request.UserID)
		if err != nil {
			response.Error = "Error deleting user"

		} else {
			response.Success = true
		}

	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. %s", err.Error())
		return
	}

	utilities.Write(conn, dataToSend)
}

func isEmailValid(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		Log.WarningF("Error email verification. %s", err.Error())
		return false
	}
	return true
}
