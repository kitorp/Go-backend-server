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
	response := library.SetQuotaResponse{}
	response.Success = false

	if Authenticate(req.Email, req.Password, req.Token,req.UserID){
		err := updateQuota(req.UserID, req.Quota)
		if len(err)>0 {
			response.Error = err
			fmt.Println(err)
		}else{
			response.Success = true
		}

	}else{
		response.Error = "Authentication Error"
	}

	fmt.Println("sending response: ", response)
	dataToSend, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	fmt.Println("response data: ", dataToSend)

	utilities.Write(conn, dataToSend)
}
