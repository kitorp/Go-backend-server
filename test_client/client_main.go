package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	connectionTest(conn)
}

func connectionTest(conn net.Conn) bool{
	text := "abcd"

	fmt.Fprintf(conn, text+"\n")

	message, _ := bufio.NewReader(conn).ReadString('\n')

	if message != "ABCD" {

		fmt.Printf("Expected: ABCD, Reply: %s",message)
		return  false

	}
	return true
}
