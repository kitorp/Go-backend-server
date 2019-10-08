package handlers

import (
	"encoding/json"
	"net"

	"../utilities"
)

func loginHandler(conn net.Conn, originalMessage []byte) {

	req := utilities.LoginRequest{}
	err := json.Unmarshal(originalMessage, &req)
	if err != nil {
		Log.WarningF("Json Unmarshal Error. %s", err.Error())
		return
	}

	response := tryLogin(req)
	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. %s", err.Error())
		return
	}

	utilities.Write(conn, dataToSend)
}
