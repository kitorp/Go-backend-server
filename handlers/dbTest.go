package handlers

import (
	"../library"
	"fmt"
)
func DBTest() {

	//tryLoginTest1()
	//authenticateByTokenTest()
	//updateQoutaTest()
	//createUserTest()
	//deleteUserTest()
	//listUserTest()
	//addResourceTest()
	deleteResourceTest()
	listResourceTest()

}

func tryLoginTest1() {
	req := library.LoginRequest{
		Email:    "protik2095@gmail.com",
		Password: "123",
	}
	resp := tryLogin(req)
	fmt.Println(resp)
}

func authenticateByTokenTest(){
	ans:= AuthenticateByToken("abcd", 1000000)
	fmt.Println("answer: ", ans)
}

func updateQoutaTest(){
	fmt.Println( updateQuota(1000000, 100) )
}

func createUserTest(){
	fmt.Println(createUser("sprotik@gmail.com", "qwe"))
}

func listUserTest(){
	fmt.Println(listUser(100,0))
}

func deleteUserTest(){
	fmt.Println(deleteUser(1000001) )
}

func addResourceTest(){
	fmt.Println(addResource(1000000, "HI there") )
}


func listResourceTest(){
	fmt.Println( listResource(1000000) )
}

func deleteResourceTest(){
	fmt.Println(deleteResource(1000000, "HI there") )
}

