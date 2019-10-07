package handlers

import (
	"../library"
	"../utilities"
	"encoding/json"
	"net"
)

func setQuotaHandler(conn net.Conn, data []byte) {

	request := library.SetQuotaRequest{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		panic(err)
	}

	response := library.CommonResponse{}
	response.Success = false

	if Authenticate(request.Email, request.Password, request.Token, request.UserID) {
		err := updateQuota(request.UserID, request.Quota)
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
