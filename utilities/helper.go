package utilities

import (
	"bufio"
	"encoding/binary"
	"github.com/apsdehal/go-logger"
	"net"
)

const (
	Login          uint32 = 1
	CreateResource uint32 = 2
	ListResource   uint32 = 3
	DeleteResource uint32 = 4
	SetQuota       uint32 = 5
	CreateUser     uint32 = 6
	ListUser       uint32 = 7
	DeleteUser     uint32 = 8

	MaximumPacketSizeInByte uint32 = 1024
)

var (
	Log    *logger.Logger
	Config Configuration
)

func EncodeMessage(messageType uint32, data []byte) (message []byte) {

	firstByte := make([]byte, 4)
	binary.BigEndian.PutUint32(firstByte, uint32(messageType))
	message = append(firstByte, data...)
	return
}

func DecodeMessage(message []byte) (messageType uint32, data []byte) {

	messageType = binary.BigEndian.Uint32(message[:4])
	data = message[4:]
	return
}

func Read(conn net.Conn) (data []byte) {

	messageInBytes := make([]byte, MaximumPacketSizeInByte)
	length, err := bufio.NewReader(conn).Read(messageInBytes)
	if err != nil {
		Log.WarningF("Read Error. ", err.Error())
		return
	}
	return messageInBytes[:length]

}

func Write(conn net.Conn, data []byte) {

	_, err := conn.Write(data)
	if err != nil {
		Log.WarningF("Write Error. ", err.Error())
		return
	}
	return
}
