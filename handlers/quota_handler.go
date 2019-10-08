package handlers

import (
	"encoding/json"
	"net"

	"../utilities"
)

func setQuotaHandler(conn net.Conn, data []byte) {

	request := utilities.SetQuotaRequest{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		Log.WarningF("Json Unmarshal Error. ", err.Error())
		return
	}

	response := utilities.CommonResponse{}
	response.Success = false

	if Authenticate(request.Email, request.Password, request.Token, request.UserID) {
		err := updateQuota(request.UserID, request.Quota)
		if err != nil {
			response.Error = err.Error()
		} else {
			response.Success = true
		}

	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. ", err.Error())
	}

	utilities.Write(conn, dataToSend)
}
