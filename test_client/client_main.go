package main

import (
	"../library"
	"../utilities"
	"encoding/json"
	"fmt"
	"net"
)

func main() {

	//loginTest()
	//createResourceTest1()
	//createResourceTest2()
	//listResourceTest1()
	listResourceTest2()

}

func loginTest() {

	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.LoginRequest{
		Email:    "Sprotik",
		Password: "123",
	}

	a, err := json.Marshal(crd)
	if err != nil {
		panic("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.Login, a)

	utilities.Write(conn,dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	receivedData := library.LoginResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		panic(err)
	}

	fmt.Println(receivedData)


}


func createResourceTest1(){
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.CreateResourceRequest{
		Email:    "Sprotik",
		Password: "123",
		Token:    "",
		Resource: "abcd",
	}

	a, err := json.Marshal(crd)
	if err != nil {
		panic("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.CreateResource, a)

	utilities.Write(conn,dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	fmt.Println("In Byte: ", receivedDatainBytes)
	receivedData := library.CreateResourceResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		panic(err)
	}

	fmt.Println(receivedData)

}

func createResourceTest2(){
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.CreateResourceRequest{
		Token:    "wow",
		Resource: "abcd",
	}

	a, err := json.Marshal(crd)
	if err != nil {
		panic("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.CreateResource, a)

	utilities.Write(conn,dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	fmt.Println("In Byte: ", receivedDatainBytes)
	receivedData := library.CreateResourceResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		panic(err)
	}

	fmt.Println(receivedData)

}

func listResourceTest1(){
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.ListResourceRequest{
		Token:    "wow",
	}

	a, err := json.Marshal(crd)
	if err != nil {
		panic("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.ListResource, a)

	utilities.Write(conn,dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	fmt.Println("In Byte: ", receivedDatainBytes)
	receivedData := library.ListResourceResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		panic(err)
	}

	fmt.Println(receivedData)
}

func listResourceTest2(){
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.ListResourceRequest{
		Email:    "Sprotik",
		Password: "123",
	}

	a, err := json.Marshal(crd)
	if err != nil {
		panic("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.ListResource, a)

	utilities.Write(conn,dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	fmt.Println("In Byte: ", receivedDatainBytes)
	receivedData := library.ListResourceResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		panic(err)
	}

	fmt.Println(receivedData)
}