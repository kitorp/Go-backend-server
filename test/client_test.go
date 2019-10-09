package main

import (
	"encoding/json"
	"net"
	"testing"

	"../utilities"
)
func getConnection()(conn net.Conn){
	conn, _ = net.Dial("tcp", "127.0.0.1:8081")
	return conn
}

func Test_Login(t *testing.T) {
	conn := getConnection()
	defer conn.Close()

	request := utilities.LoginRequest{
		Email:    "sportik@gmail.com",
		Password: "qwe",
	}

	requestInByte, err := json.Marshal(request)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.Login, requestInByte)

	utilities.Write(conn, dataToSend)

	receivedDataInBytes := utilities.Read(conn)

	receivedData := utilities.LoginResponse{}

	err = json.Unmarshal(receivedDataInBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}
	if !receivedData.Success {
		t.Error(receivedData)
	}
}

func Test_createResource(t *testing.T) {

	conn := getConnection()
	defer conn.Close()

	request := utilities.CreateResourceRequest{
		Email:    "p95@gmail.com",
		Password: "123",
		Token:    "",
		UserID:   1000000,
		Resource: "my resource",
	}

	requestInBytes, err := json.Marshal(request)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.CreateResource, requestInBytes)

	utilities.Write(conn, dataToSend)

	receivedDataInBytes := utilities.Read(conn)

	receivedData := utilities.CommonResponse{}

	err = json.Unmarshal(receivedDataInBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}
	if receivedData.Success {
		t.Error(receivedData)
	}
}

func Test_listResource(t *testing.T) {
	conn := getConnection()
	defer conn.Close()

	request := utilities.ListResourceRequest{
		Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InByb3RpazIwOTVAZ21haWwuY29tIiwiZXhwIjoxNTcwNTQ5NzU5fQ.324xoWAPu0LHXNHYNIKkk6VVSAqG20Vy794zh-iuQkM",
		UserID: 1000000,
	}

	requestInBytes, err := json.Marshal(request)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.ListResource, requestInBytes)

	utilities.Write(conn, dataToSend)

	receivedDataInBytes := utilities.Read(conn)

	receivedData := utilities.ListResourceResponse{}

	err = json.Unmarshal(receivedDataInBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}

}

func Test_listResource2(t *testing.T) {
	conn := getConnection()
	defer conn.Close()

	request := utilities.ListResourceRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		UserID:   1000001,
	}

	requestInBytes, err := json.Marshal(request)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.ListResource, requestInBytes)

	utilities.Write(conn, dataToSend)

	receivedDataInBytes := utilities.Read(conn)

	receivedData := utilities.ListResourceResponse{}

	err = json.Unmarshal(receivedDataInBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}
}

func Test_deleteResource(t *testing.T) {
	conn := getConnection()
	defer conn.Close()

	request := utilities.DeleteResourceRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		UserID:   1000001,
		Resource: "should work",
	}

	requestInBytes, err := json.Marshal(request)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.DeleteResource, requestInBytes)

	utilities.Write(conn, dataToSend)

	receivedDataInBytes := utilities.Read(conn)

	receivedData := utilities.CommonResponse{}

	err = json.Unmarshal(receivedDataInBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}
	if !receivedData.Success {
		t.Error(receivedData)
	}
}

func Test_setQouta(t *testing.T) {
	conn := getConnection()
	defer conn.Close()

	request := utilities.SetQuotaRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		UserID:   1000001,
		Quota:    10,
	}

	requestInBytes, err := json.Marshal(request)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.SetQuota, requestInBytes)

	utilities.Write(conn, dataToSend)

	receivedDataInBytes := utilities.Read(conn)

	receivedData := utilities.CommonResponse{}

	err = json.Unmarshal(receivedDataInBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}
	if !receivedData.Success {
		t.Error(receivedData)
	}
}

func Test_createUser(t *testing.T) {

	conn := getConnection()
	defer conn.Close()

	request := utilities.CreateUserRequest{
		Email:        "protik2095@gmail.com",
		Password:     "123",
		UserEmail:    "sportik@gmail.com",
		UserPassword: "qwe",
		UserType:     0,
	}

	requestInBytes, err := json.Marshal(request)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.CreateUser, requestInBytes)

	utilities.Write(conn, dataToSend)

	receivedDataInBytes := utilities.Read(conn)

	receivedData := utilities.CommonResponse{}

	err = json.Unmarshal(receivedDataInBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}

	if receivedData.Success {
		t.Error(receivedData)
	}

}

func Test_listUser(t *testing.T) {

	conn := getConnection()
	defer conn.Close()

	request := utilities.ListUserRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		Limit:    2,
		Offset:   1,
	}

	requestInBytes, err := json.Marshal(request)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.ListUser, requestInBytes)

	utilities.Write(conn, dataToSend)

	receivedDataInBytes := utilities.Read(conn)

	receivedData := utilities.ListUserResponse{}

	err = json.Unmarshal(receivedDataInBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}
	if !receivedData.Success {
		t.Error(receivedData)
	}
}

func Test_deleteUserResponse(t *testing.T) {
	conn := getConnection()
	defer conn.Close()

	request := utilities.DeleteUserRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		UserID:   1000002,
	}

	requestInByte, err := json.Marshal(request)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.DeleteUser, requestInByte)

	utilities.Write(conn, dataToSend)

	receivedDataInBytes := utilities.Read(conn)

	receivedData := utilities.CommonResponse{}

	err = json.Unmarshal(receivedDataInBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}
	if !receivedData.Success {
		t.Error(receivedData)
	}
}
