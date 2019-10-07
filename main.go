package main

import (
	"./handlers"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting Server")
	handlers.DB = handlers.GetDB()
	//handlers.DBTest()

	//return //remove it
	listen, _ := net.Listen("tcp", ":8081")
	defer listen.Close()

	for {
		conn, _ := listen.Accept()
		defer conn.Close()
		handlers.Handler(conn)
	}
}

