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
		Log.WarningF("Json Unmarshal Error. %s", err.Error())
		return
	}

	response := utilities.CommonResponse{
		Success: false,
	}
	if authenticateUser(request.Email, request.Password, request.Token, request.UserID) {
		if !canAddResource(request.UserID) {
			response.Error = "Quota limit exceeded"
		} else {
			err := addResource(request.UserID, request.Resource)
			if err != nil {
				response.Error = "Error creating resource"
			} else {
				response.Success = true
			}
		}

	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. %s", err.Error())
		return
	}

	utilities.Write(conn, dataToSend)
}

func listResourceHandler(conn net.Conn, originalMessage []byte) {

	request := utilities.ListResourceRequest{}
	err := json.Unmarshal(originalMessage, &request)
	if err != nil {
		Log.WarningF("Json Unmarshal Error. %s", err.Error())
		return
	}

	response := utilities.ListResourceResponse{}

	if authenticateUser(request.Email, request.Password, request.Token, request.UserID) {
		list, err := listResource(request.UserID)
		if err != nil {
			response.Error = "Error listing resource"

		} else {
			response.Resource = append(response.Resource, list...)
		}

	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. %s", err.Error())
		return
	}

	utilities.Write(conn, dataToSend)
}

func deleteResourceHandler(conn net.Conn, originalMessage []byte) {

	request := utilities.DeleteResourceRequest{}
	err := json.Unmarshal(originalMessage, &request)
	if err != nil {
		Log.WarningF("Json Unmarshal Error. %s", err.Error())
		return
	}

	response := utilities.CommonResponse{
		Success: false,
	}

	if authenticateUser(request.Email, request.Password, request.Token, request.UserID) {
		err := deleteResource(request.UserID, request.Resource)
		if err != nil {
			response.Error = "Error deleting resource"
		} else {
			response.Success = true
		}
	} else {
		response.Error = "Authentication Error"
	}

	dataToSend, err := json.Marshal(response)
	if err != nil {
		Log.WarningF("Json Marshal Error. %s", err.Error())
		return
	}

	utilities.Write(conn, dataToSend)
}


func addResource(userId int, resource string) (err error) {

	list, err := listResource(userId)
	if err != nil {
		return err
	}
	list = append(list, resource)

	err = modifyResource(userId, list)

	if err != nil {
		return err
	}

	err = addResourceCount(userId, 1)
	if err != nil {
		return err
	}
	return nil
}

func deleteResource(userId int, resource string) (err error) {

	list, err := listResource(userId)
	if err != nil {
		return
	}

	var newList []string
	numberOfDeletedResource := 0
	for _, s := range list {
		if s == resource {
			numberOfDeletedResource++
			continue
		}
		newList = append(newList, s)
	}

	err = modifyResource(userId, newList)
	if err != nil {
		return
	}

	err = addResourceCount(userId, -numberOfDeletedResource)
	if err != nil {
		return
	}
	return nil
}