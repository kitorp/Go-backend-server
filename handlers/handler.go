package handlers

import (
	"../utilities"
	"fmt"
	"net"
)

func Handler(conn net.Conn) {
	messageInBytes := utilities.Read(conn)
	messageType, originalMessage := utilities.DecodeMessage(messageInBytes)
	fmt.Println("original message: ", originalMessage)

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
