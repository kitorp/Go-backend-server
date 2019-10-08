package main

import (
	"./handlers"
	"fmt"
	"github.com/apsdehal/go-logger"
	"net"
	"os"
)

func main() {

	fmt.Println("Starting Server")

	f, err := os.OpenFile("./logs/server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer f.Close()
	log, err := logger.New("backend-server", 1, f)
	if err != nil {
		panic("Error initiating log")
	}


	handlers.Log = log
	handlers.DB = handlers.GetDB()

	listen, _ := net.Listen("tcp", ":8081")
	defer listen.Close()
	for {
		conn, _ := listen.Accept()
		defer conn.Close()
		handlers.Handler(conn)
	}
}
