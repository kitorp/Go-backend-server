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
	response := library.CreateUserResponse{}
	if Authenticate(req.Email, req.Password, req.Token,0){
		err := createUser(req.UserEmail, req.UserPassword)
		if len(err)>0 {
			response.Error = err
			fmt.Println(err)
		}else{
			response.Success = true
		}

	}else{
		response.Error = "Authentication Error"
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
		Success:false,
	}


	if Authenticate(req.Email, req.Password, req.Token,0){
		list , err := listUser(req.Limit, req.Offset)
		if len(err)>0 {
			response.Error = err
			fmt.Println(err)
		}else{
			response.Users = append(response.Users, list...)
			response.Success = true
		}

	}else{
		response.Error = "Authentication Error"
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
		Success:false,
	}
	if Authenticate(req.Email, req.Password, req.Token,0){
		err := deleteUser(req.UserID)
		if len(err)>0 {
			response.Error = err
			fmt.Println(err)
		}else{
			response.Success = true
		}

	}else{
		response.Error = "Authentication Error"
	}
	fmt.Println("sending response: ", response)
	dataToSend, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	fmt.Println("response data: ", dataToSend)

	utilities.Write(conn, dataToSend)
}
