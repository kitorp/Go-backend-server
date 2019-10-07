package utilities

import (
	"encoding/binary"
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
