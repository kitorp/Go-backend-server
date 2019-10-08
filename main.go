package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"./handlers"
	"./utilities"

	"github.com/apsdehal/go-logger"
)

func readConfig() utilities.Configuration {

	jsonFile, err := os.Open("./config.json")
	if err != nil {
		fmt.Println("Error opening config file")
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading config file")
	}
	var newConfig utilities.Configuration
	err = json.Unmarshal(byteValue, &newConfig)
	if err != nil {
		fmt.Println(err)
	}
	return newConfig
}

func main() {

	fmt.Println("Starting Server")

	utilities.Config = readConfig()
	file, err := os.OpenFile(utilities.Config.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file: ", err)
	}
	defer file.Close()
	log, err := logger.New("backend-server", 1, file)
	if err != nil {
		fmt.Println("Error initiating log. Error: ", err)
	}

	utilities.Log = log
	handlers.Log = handlers.GetLogger()
	handlers.DB = handlers.GetDB()

	listen, _ := net.Listen(utilities.Config.ConnType, utilities.Config.ConnAddress+":"+utilities.Config.ConnPort)
	defer listen.Close()
	for {
		conn, _ := listen.Accept()
		defer conn.Close()
		go handlers.Handler(conn)
	}
}
