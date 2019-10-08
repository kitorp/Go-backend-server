package handlers

import (
	"net"

	"../utilities"
	"github.com/apsdehal/go-logger"
)

const (
	passwordLength = 1
	emailLength    = 1
	tokenLength    = 1

	userAdmin   = 1
	userGeneral = 0
)

var (
	Log *logger.Logger
)

func GetLogger() *logger.Logger {
	return utilities.Log
}

func Handler(conn net.Conn) {

	messageInBytes := utilities.Read(conn)
	messageType, data := utilities.DecodeMessage(messageInBytes)

	Log.InfoF("Request type: %d", messageType)

	if messageType == utilities.Login {
		loginHandler(conn, data)
	} else if messageType == utilities.CreateResource {
		createResourceHandler(conn, data)
	} else if messageType == utilities.ListResource {
		listResourceHandler(conn, data)
	} else if messageType == utilities.DeleteResource {
		deleteResourceHandler(conn, data)
	} else if messageType == utilities.SetQuota {
		setQuotaHandler(conn, data)
	} else if messageType == utilities.CreateUser {
		createUserHandler(conn, data)
	} else if messageType == utilities.ListUser {
		listUserHandler(conn, data)
	} else if messageType == utilities.DeleteUser {
		deleteUserHandler(conn, data)
	} else {
		return
	}
}


