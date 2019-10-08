package handlers

import (
	"../library"
	"../utilities"
	"encoding/json"
	"net"
)

func createUserHandler(conn net.Conn, data []byte) {

	request := library.CreateUserRequest{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		Log.WarningF("Json Unmarshal Error. ", err.Error())
		return
	}

	response := library.CommonResponse{
		Success: false,
	}

	if Authenticate(request.Email, request.Password, request.Token, 0) {
		err := createUser(request.UserEmail, request.UserPassword, request.UserType)
		if err != nil {
			response.Error = err.Error()
		} else {
			response.Success = true
		}

	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. ", err.Error())
		return
	}

	utilities.Write(conn, dataToSend)
}

func listUserHandler(conn net.Conn, data []byte) {

	request := library.ListUserRequest{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		Log.WarningF("Json Unmarshal Error. ", err.Error())
		return
	}

	response := library.ListUserResponse{
		Success: false,
	}

	if Authenticate(request.Email, request.Password, request.Token, 0) {
		list, err := listUser(request.Limit, request.Offset)
		if err != nil {
			response.Error = err.Error()
		} else {
			response.Users = append(response.Users, list...)
			response.Success = true
		}

	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. ", err.Error())
		return
	}

	utilities.Write(conn, dataToSend)
}

func deleteUserHandler(conn net.Conn, data []byte) {

	request := library.DeleteUserRequest{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		Log.WarningF("Json Unmarshal Error. ", err.Error())
		return
	}

	response := library.CommonResponse{
		Success: false,
	}
	if Authenticate(request.Email, request.Password, request.Token, 0) {
		err := deleteUser(request.UserID)
		if err != nil {
			response.Error = err.Error()

		} else {
			response.Success = true
		}

	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. ", err.Error())
		return
	}

	utilities.Write(conn, dataToSend)
}
