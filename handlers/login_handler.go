package handlers

import (
	"../library"
	"../utilities"
	"encoding/json"
	"fmt"
	"net"
)

func loginHandler(conn net.Conn, originalMessage []byte) {

	req := library.LoginRequest{}
	err := json.Unmarshal(originalMessage, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println("req ", req)
	response := tryLogin(req)
	fmt.Println("sending response: ", response)
	dataToSend, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	fmt.Println("response data: ", dataToSend)

	utilities.Write(conn, dataToSend)
}
