package handlers

import (
	"../utilities"
	"fmt"
	"net"
)

const(
	passwordLength = 1
	emailLength = 1
	tokenLength = 1


	userAdmin = 1
	userGeneral = 0
)

func Handler(conn net.Conn) {
	messageInBytes := utilities.Read(conn)
	messageType, originalMessage := utilities.DecodeMessage(messageInBytes)
	fmt.Println("message type: ", messageType)

	if messageType == utilities.Login {
		loginHandler(conn, originalMessage)
	} else if messageType == utilities.CreateResource {
		createResourceHandler(conn, originalMessage)
	} else if messageType == utilities.ListResource {
		listResourceHandler(conn, originalMessage)
	} else if messageType == utilities.DeleteResource {
		deleteResourceHandler(conn, originalMessage)
	} else if messageType == utilities.SetQuota {
		setQuotaHandler(conn, originalMessage)
	} else if messageType == utilities.CreateUser{
		createUserHandler(conn, originalMessage)
	} else if messageType == utilities.ListUser{
		listUserHandler(conn, originalMessage)
	} else if messageType == utilities.DeleteUser {
		deleteUserHandler(conn, originalMessage)
	}
}

func Authenticate(email string, password string, token string, userid int) bool {
	fmt.Println("Start to Authenticate")
	if len(token)>= tokenLength {
		fmt.Println("here ?")
		return AuthenticateByToken(token, userid)
	}else if len(email)>= emailLength && len(password)>=passwordLength {
		return AuthenticateByEmailPassword(email, password, userid)
	}else{
		return false

	}

}