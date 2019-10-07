package utilities

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
)

func Read(conn net.Conn) (data []byte) {

	messageInBytes := make([]byte, 1024)
	_ ,err := bufio.NewReader(conn).Read(messageInBytes)
	if err != nil {
		fmt.Println(err)
		panic(" bufio read err")
	}
	length := binary.BigEndian.Uint32(messageInBytes[:4])
	return messageInBytes[4:length+4]

}

func Write(conn net.Conn,data []byte){

	length := len(data)
	firstByte := make([]byte, 4)
	binary.BigEndian.PutUint32(firstByte, uint32(length))

	newbyte := append(firstByte,data...)
	_, err := conn.Write(newbyte)
	if err != nil {
		panic("data write error")
	}
}

