package handlers

import (
	"../library"
	"../utilities"
	"github.com/apsdehal/go-logger"
	"net"
)

const (
	passwordLength = 1
	emailLength    = 1
	tokenLength    = 1

	userAdmin   = 1
	userGeneral = 0
)

var (
	Log    *logger.Logger
	Config library.Configuration
)

func GetLogger() *logger.Logger {
	return Log
}

func Handler(conn net.Conn) {

	log := GetLogger()

	messageInBytes := utilities.Read(conn)
	messageType, data := utilities.DecodeMessage(messageInBytes)

	log.InfoF("Request type: %d", messageType)

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

func Authenticate(email string, password string, token string, userid int) bool {

	if len(token) >= tokenLength {
		return AuthenticateByToken(token, userid)
	} else if len(email) >= emailLength && len(password) >= passwordLength {
		return AuthenticateByEmailPassword(email, password, userid)
	} else {
		return false
	}

}
