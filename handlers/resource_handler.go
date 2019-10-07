package handlers

import (
	"../library"
	"../utilities"
	"encoding/json"
	"fmt"
	"net"
)

func createResourceHandler(conn net.Conn, originalMessage []byte) {
	req := library.CreateResourceRequest{}
	err := json.Unmarshal(originalMessage, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println("req ", req)
	response := library.CreateResourceResponse{
		Success: false,
	}
	if Authenticate(req.Email, req.Password, req.Token, req.UserID) {
		if !canAddResource(req.UserID) {
			response.Error = "Qouta limit exceeded"
		} else {
			err := addResource(req.UserID, req.Resource)
			if len(err) > 0 {
				response.Error = err
				fmt.Println(err)
			} else {
				response.Success = true
			}
		}

	} else {
		response.Error = "Authentication Error"
	}
	fmt.Println("sending response: ", response)
	dataToSend, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	utilities.Write(conn, dataToSend)
}

func listResourceHandler(conn net.Conn, originalMessage []byte) {
	req := library.ListResourceRequest{}
	err := json.Unmarshal(originalMessage, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println("req ", req)
	response := library.ListResourceResponse{}

	if Authenticate(req.Email, req.Password, req.Token, req.UserID) {
		list, err := listResource(req.UserID)
		if len(err) > 0 {
			response.Error = err
			fmt.Println(err)
		} else {
			response.Resource = append(response.Resource, list...)
		}

	} else {
		response.Error = "Authentication Error"
	}

	fmt.Println("sending response: ", response)
	dataToSend, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	utilities.Write(conn, dataToSend)
}

func deleteResourceHandler(conn net.Conn, originalMessage []byte) {
	req := library.DeleteResourceRequest{}
	err := json.Unmarshal(originalMessage, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println("req ", req)
	response := library.DeleteResourceResponse{
		Success: false,
	}
	if Authenticate(req.Email, req.Password, req.Token, req.UserID) {
		err := deleteResource(req.UserID, req.Resource)
		if len(err) > 0 {
			response.Error = err
			fmt.Println(err)
		} else {
			response.Success = true
		}

	} else {
		response.Error = "Authentication Error"
	}
	fmt.Println("sending response: ", response)
	dataToSend, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	utilities.Write(conn, dataToSend)
}
