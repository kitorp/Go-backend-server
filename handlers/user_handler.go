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
		panic(err)
	}

	response := library.CommonResponse{
		Success: false,
	}

	if Authenticate(request.Email, request.Password, request.Token, 0) {
		err := createUser(request.UserEmail, request.UserPassword)
		if len(err) > 0 {
			response.Error = err
		} else {
			response.Success = true
		}

	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	utilities.Write(conn, dataToSend)
}

func listUserHandler(conn net.Conn, data []byte) {

	request := library.ListUserRequest{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		panic(err)
	}

	response := library.ListUserResponse{
		Success: false,
	}

	if Authenticate(request.Email, request.Password, request.Token, 0) {
		list, err := listUser(request.Limit, request.Offset)
		if len(err) > 0 {
			response.Error = err
		} else {
			response.Users = append(response.Users, list...)
			response.Success = true
		}

	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	utilities.Write(conn, dataToSend)
}

func deleteUserHandler(conn net.Conn, data []byte) {

	request := library.DeleteUserRequest{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		panic(err)
	}

	response := library.CommonResponse{
		Success: false,
	}
	if Authenticate(request.Email, request.Password, request.Token, 0) {
		err := deleteUser(request.UserID)
		if len(err) > 0 {
			response.Error = err

		} else {
			response.Success = true
		}

	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	utilities.Write(conn, dataToSend)
}
