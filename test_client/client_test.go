package main

import (
	"encoding/json"
	"net"
	"testing"

	"../library"
	"../utilities"
)

func Test_Login(t *testing.T) {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.LoginRequest{
		Email:    "sportik@gmail.com",
		Password: "qwe",
	}

	a, err := json.Marshal(crd)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.Login, a)

	utilities.Write(conn, dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	receivedData := library.LoginResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}
	if !receivedData.Success {
		t.Error("Failed")
	}
}

func Test_createResource(t *testing.T) {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.CreateResourceRequest{
		Email:    "p95@gmail.com",
		Password: "123",
		Token:    "",
		UserID:   1000000,
		Resource: "my resource",
	}

	a, err := json.Marshal(crd)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.CreateResource, a)

	utilities.Write(conn, dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	receivedData := library.CommonResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}
	if receivedData.Success {
		t.Error("Failed")
	}
}

func Test_listResource(t *testing.T) {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.ListResourceRequest{
		Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InByb3RpazIwOTVAZ21haWwuY29tIiwiZXhwIjoxNTcwNTQ5NzU5fQ.324xoWAPu0LHXNHYNIKkk6VVSAqG20Vy794zh-iuQkM",
		UserID: 1000000,
	}

	a, err := json.Marshal(crd)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.ListResource, a)

	utilities.Write(conn, dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	receivedData := library.ListResourceResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}

}

func Test_listResource2(t *testing.T) {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.ListResourceRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		UserID:   1000001,
	}

	a, err := json.Marshal(crd)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.ListResource, a)

	utilities.Write(conn, dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	receivedData := library.ListResourceResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}
}

func Test_deleteResource(t *testing.T) {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.DeleteResourceRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		UserID:   1000001,
		Resource: "should work",
	}

	a, err := json.Marshal(crd)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.DeleteResource, a)

	utilities.Write(conn, dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	receivedData := library.CommonResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}
	if !receivedData.Success {
		t.Error("Failed")
	}
}

func Test_setQouta(t *testing.T) {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.SetQuotaRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		UserID:   1000001,
		Quota:    10,
	}

	a, err := json.Marshal(crd)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.SetQuota, a)

	utilities.Write(conn, dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	receivedData := library.CommonResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}
	if !receivedData.Success {
		t.Error("Failed")
	}
}

func Test_createUser(t *testing.T) {

	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.CreateUserRequest{
		Email:        "protik2095@gmail.com",
		Password:     "123",
		UserEmail:    "sportik@gmail.com",
		UserPassword: "qwe",
		UserType:     0,
	}

	a, err := json.Marshal(crd)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.CreateUser, a)

	utilities.Write(conn, dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	receivedData := library.CommonResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}

	if receivedData.Success {
		t.Error(receivedData.Error)
	}

}

func Test_listUser(t *testing.T) {

	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.ListUserRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		Limit:    2,
		Offset:   1,
	}

	a, err := json.Marshal(crd)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.ListUser, a)

	utilities.Write(conn, dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	receivedData := library.ListUserResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}
	if !receivedData.Success {
		t.Error("Failed")
	}
}

func Test_deleteUserResponse(t *testing.T) {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	crd := library.DeleteUserRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
		UserID:   1000002,
	}

	a, err := json.Marshal(crd)
	if err != nil {
		t.Error("json marshal error")
	}
	dataToSend := utilities.EncodeMessage(utilities.DeleteUser, a)

	utilities.Write(conn, dataToSend)

	receivedDatainBytes := utilities.Read(conn)

	receivedData := library.CommonResponse{}

	err = json.Unmarshal(receivedDatainBytes, &receivedData)
	if err != nil {
		t.Error(err)
	}
	if !receivedData.Success {
		t.Error("Failed")
	}
}
