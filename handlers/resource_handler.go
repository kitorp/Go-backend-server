package handlers

import (
	"encoding/json"
	"net"

	"../utilities"
)

func createResourceHandler(conn net.Conn, data []byte) {

	request := utilities.CreateResourceRequest{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		Log.WarningF("Json Unmarshal Error. ", err.Error())
		return
	}

	response := utilities.CommonResponse{
		Success: false,
	}
	if Authenticate(request.Email, request.Password, request.Token, request.UserID) {
		if !canAddResource(request.UserID) {
			response.Error = "Qouta limit exceeded"
		} else {
			err := addResource(request.UserID, request.Resource)
			if err != nil {
				response.Error = err.Error()
			} else {
				response.Success = true
			}
		}

	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. ", err.Error())
		return
	}

	utilities.Write(conn, dataToSend)
}

func listResourceHandler(conn net.Conn, originalMessage []byte) {
	request := utilities.ListResourceRequest{}
	err := json.Unmarshal(originalMessage, &request)
	if err != nil {
		Log.WarningF("Json Unmarshal Error. ", err.Error())
		return
	}

	response := utilities.ListResourceResponse{}

	if Authenticate(request.Email, request.Password, request.Token, request.UserID) {
		list, err := listResource(request.UserID)
		if err != nil {
			response.Error = err.Error()

		} else {
			response.Resource = append(response.Resource, list...)
		}

	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. ", err.Error())
		return
	}

	utilities.Write(conn, dataToSend)
}

func deleteResourceHandler(conn net.Conn, originalMessage []byte) {
	request := utilities.DeleteResourceRequest{}
	err := json.Unmarshal(originalMessage, &request)
	if err != nil {
		Log.WarningF("Json Unmarshal Error. ", err.Error())
		return
	}

	response := utilities.CommonResponse{
		Success: false,
	}
	if Authenticate(request.Email, request.Password, request.Token, request.UserID) {
		err := deleteResource(request.UserID, request.Resource)
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
		return
	}

	utilities.Write(conn, dataToSend)
}
