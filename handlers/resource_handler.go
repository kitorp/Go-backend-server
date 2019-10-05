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
		Email:    req.Email,
		Password: req.Password,
		Token:    req.Token,
		Resource: req.Resource,
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

func listResourceHandler(conn net.Conn, originalMessage []byte) {
	req := library.ListResourceRequest{}
	err := json.Unmarshal(originalMessage, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println("req ", req)
	response := library.ListResourceResponse{
		Email:    req.Email,
		Password: req.Password,
		Token:    req.Token,
	}
	var all []string
	all = append(all, "abc","dfs", "dfsf")
	if req.Email == "Sprotik" && req.Password == "123" {
		fmt.Println("correct if")
		response.Resource = append(response.Resource, all...)

	} else if req.Token == "wow" {
		response.Resource = append(response.Resource, all...)
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

func deleteResourceHandler(conn net.Conn, originalMessage []byte) {
	req := library.DeleteResourceRequest{}
	err := json.Unmarshal(originalMessage, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println("req ", req)
	response := library.DeleteResourceResponse{
		Email:    req.Email,
		Password: req.Password,
		Token:    req.Token,
		Resource: req.Resource,
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

