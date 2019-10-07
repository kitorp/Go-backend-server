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
	listResourceTest1()
	//listResourceTest2()

	//deleteResourceTest()
	//setQoutaTest()
	//createUserTest()
	//listUserTest()
	//deleteUserResponseTest()



}

func loginTest() {

	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.LoginRequest{
		Email:    "protik2095@gmail.com",
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
		Email:    "sprotik@gmail.com",
		Password: "qwe",
		Token:    "",
		UserID: 1000000,
		Resource: "your life not my rules",
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
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InByb3RpazIwOTVAZ21haWwuY29tIiwiZXhwIjoxNTcwNTQ5NzU5fQ.324xoWAPu0LHXNHYNIKkk6VVSAqG20Vy794zh-iuQkM",
		UserID: 1000000,
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
		Email:    "protik2095@gmail.com",
		Password: "123",
		UserID: 1000001,
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


func deleteResourceTest(){
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.DeleteResourceRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		UserID: 1000001,
		Resource:"should work",
	}

	a, err := json.Marshal(crd)
	if err != nil {
		panic("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.DeleteResource, a)

	utilities.Write(conn,dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	fmt.Println("In Byte: ", receivedDatainBytes)
	receivedData := library.DeleteResourceResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		panic(err)
	}

	fmt.Println(receivedData)
}



func setQoutaTest(){
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.SetQuotaRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		UserID: 1000001,
		Quota: 10,
	}

	a, err := json.Marshal(crd)
	if err != nil {
		panic("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.SetQuota, a)

	utilities.Write(conn,dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	fmt.Println("In Byte: ", receivedDatainBytes)
	receivedData := library.SetQuotaResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		panic(err)
	}

	fmt.Println(receivedData)
}


func createUserTest(){
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.CreateUserRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		UserEmail: "arif@gmail.com",
		UserPassword:"baba",
	}

	a, err := json.Marshal(crd)
	if err != nil {
		panic("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.CreateUser, a)

	utilities.Write(conn,dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	fmt.Println("In Byte: ", receivedDatainBytes)
	receivedData := library.CreateUserResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		panic(err)
	}

	fmt.Println(receivedData)

}

func listUserTest(){
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.ListUserRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		Limit:2,
		Offset:1,
	}

	a, err := json.Marshal(crd)
	if err != nil {
		panic("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.ListUser, a)

	utilities.Write(conn,dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	fmt.Println("In Byte: ", receivedDatainBytes)
	receivedData := library.ListUserResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		panic(err)
	}

	fmt.Println(receivedData)
}


func deleteUserResponseTest(){
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.DeleteUserRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		UserID:1000002,
	}

	a, err := json.Marshal(crd)
	if err != nil {
		panic("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.DeleteUser, a)

	utilities.Write(conn,dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	fmt.Println("In Byte: ", receivedDatainBytes)
	receivedData := library.DeleteUserResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		panic(err)
	}

	fmt.Println(receivedData)
}