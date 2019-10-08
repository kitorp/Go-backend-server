package main

import (
	"./handlers"
	"./library"
	"encoding/json"
	"fmt"
	"github.com/apsdehal/go-logger"
	"io/ioutil"
	"net"
	"os"
)

func readConfig() library.Configuration {

	jsonFile, err := os.Open("./config.json")
	if err != nil {
		fmt.Println("Error opening config file")
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading config file")
	}
	var newConfig library.Configuration
	err = json.Unmarshal(byteValue, &newConfig)
	if err != nil {
		fmt.Println(err)
	}
	return newConfig
}

func main() {

	fmt.Println("Starting Server")

	handlers.Config =  readConfig()

	f, err := os.OpenFile(handlers.Config.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file: ", err)
	}
	defer f.Close()
	log, err := logger.New("backend-server", 1, f)
	if err != nil {
		fmt.Println("Error initiating log. Error: ",err)
	}

	handlers.Log = log
	handlers.DB = handlers.GetDB()

	listen, _ := net.Listen(handlers.Config.ConnType, handlers.Config.ConnAddress+":"+ handlers.Config.ConnPort)
	defer listen.Close()
	for {
		conn, _ := listen.Accept()
		defer conn.Close()
		go handlers.Handler(conn)
	}
}
