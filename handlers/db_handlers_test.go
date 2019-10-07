package handlers

import "fmt"

func tryLoginTest1(){
	req := library.LoginRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
	}
	resp := tryLogin(req)
	fmt.Println(resp)
}
