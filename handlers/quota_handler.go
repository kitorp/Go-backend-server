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
		Log.WarningF("Json Unmarshal Error. %s", err.Error())
		return
	}

	response := utilities.CommonResponse{}
	response.Success = false

	if authenticateUser(request.Email, request.Password, request.Token, request.UserID) {
		err := updateQuota(request.UserID, request.Quota)
		if err != nil {
			response.Error = "Error updating Quota"
		} else {
			response.Success = true
		}
	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. %s", err.Error())
	}

	utilities.Write(conn, dataToSend)
}
