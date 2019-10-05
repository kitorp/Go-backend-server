package handlers

import (
	"../library"
	"../utilities"
	"encoding/json"
	"fmt"
	"net"
)

func createUserHandler(conn net.Conn, originalMessage []byte) {
	req := library.CreateUserRequest{}
	err := json.Unmarshal(originalMessage, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println("req ", req)
	response := library.CreateUserResponse{
		Email:        req.Email,
		Password:     req.Password,
		Token:        req.Token,
		UserEmail:    req.UserEmail,
		UserPassword: req.UserPassword,
	}
	if req.Email == "Sprotik" && req.Password == "123" {
		fmt.Println("correct if")
		response.UserID = 999
		response.Success = true

	} else if req.Token == "wow" {
		response.UserID = 999
		response.Success = true
	} else {
		response.Success = false
		response.Error = "authentication error"
	}
	fmt.Println("sending response: ", response)
	dataToSend, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	fmt.Println("response data: ", dataToSend)

	utilities.Write(conn, dataToSend)
}

func listUserHandler(conn net.Conn, originalMessage []byte) {
	req := library.ListUserRequest{}
	err := json.Unmarshal(originalMessage, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println("req ", req)
	response := library.ListUserResponse{
		Email:    req.Email,
		Password: req.Password,
		Token:    req.Token,
	}
	var all []library.User
	if req.Email == "Sprotik" && req.Password == "123" {
		fmt.Println("correct if")
		response.Users = append(response.Users, all...)

	} else if req.Token == "wow" {
		response.Users = append(response.Users, all...)
	} else {
		response.Error = "authentication error"
	}
	fmt.Println("sending response: ", response)
	dataToSend, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	fmt.Println("response data: ", dataToSend)

	utilities.Write(conn, dataToSend)
}

func deleteUserHandler(conn net.Conn, originalMessage []byte) {
	req := library.DeleteUserRequest{}
	err := json.Unmarshal(originalMessage, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println("req ", req)
	response := library.DeleteUserResponse{
		Email:    req.Email,
		Password: req.Password,
		Token:    req.Token,
		UserID:   req.UserID,
	}
	if req.Email == "Sprotik" && req.Password == "123" {
		fmt.Println("correct if")
		response.Success = true

	} else if req.Token == "wow" {
		response.Success = true
	} else {
		response.Success = false
		response.Error = "authentication error"
	}
	fmt.Println("sending response: ", response)
	dataToSend, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	fmt.Println("response data: ", dataToSend)

	utilities.Write(conn, dataToSend)
}
