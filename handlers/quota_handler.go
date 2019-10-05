package handlers

import (
	"../library"
	"../utilities"
	"encoding/json"
	"fmt"
	"net"
)

func setQuotaHandler(conn net.Conn, originalMessage []byte) {
	req := library.SetQuotaRequest{}
	err := json.Unmarshal(originalMessage, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println("req ", req)
	response := library.SetQuotaResponse{
		Email:    req.Email,
		Password: req.Password,
		Token:    req.Token,
		UserID: req.UserID,
		Quota: req.Quota,
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
