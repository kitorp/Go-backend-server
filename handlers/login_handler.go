package handlers

import (
	"../library"
	"../utilities"
	"encoding/json"
	"net"
)

func loginHandler(conn net.Conn, originalMessage []byte) {

	req := library.LoginRequest{}
	err := json.Unmarshal(originalMessage, &req)
	if err != nil {
		panic(err)
	}

	response := tryLogin(req)
	dataToSend, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	utilities.Write(conn, dataToSend)
}
