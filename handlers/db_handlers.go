package handlers

import (
	"../library"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)


func GetDB() *sql.DB {
	Db, err := sql.Open("mysql", "protik:123@/users")
	if err != nil {
		fmt.Println(err.Error())
	}
	return Db
}

func tryLogin(request library.LoginRequest) (ret library.LoginResponse) {
	fmt.Printf("SELECT password, userid FROM user_information WHERE email = %s and deleted = 0\n ", request.Email)

	row, err := DB.Query("SELECT password, userid FROM user_information WHERE email = ? and deleted = 0", request.Email)
	if err != nil {
		panic(err)
	}
	var L string
	var id int

	for row.Next() {

		err = row.Scan(&L, &id)
		if err != nil {
			panic(err)
		}
	}

	if passwordMatch(request.Password, L) {
		token, err := getToken(request.Email)
		if err != nil {
			ret.Error = err.Error()
			return
		}
		ret.UserID = id
		ret.Token = token
		return
	} else {
		ret.Error = "Wrong email/password"
		return
	}

}

func passwordMatch(givenPassword string, hashedPassword string) bool {
	hashedValue := getHashValue(givenPassword)
	return hashedValue == hashedPassword
}

func getHashValue(value string) string {
	return value
}

func getToken(email string) (token string, err error) {
	token = generateToken()
	_, err = DB.Exec("update user_information set token = ? where email = ?", token, email)
	if err != nil {
		panic(err)
	}
	return token, nil
}

func generateToken() string {
	return "abcd"
}

func AuthenticateByToken(token string, userid int) bool {
	row, err := DB.Query("SELECT userid, usertype FROM user_information WHERE token = ? and deleted = 0", token)
	if err != nil {
		panic(err)
	}
	var id, userType int

	for row.Next() {

		err = row.Scan(&id, &userType)
		if err != nil {
			panic(err)
		}
	}
	if userType == userAdmin || id == userid {
		return true
	}
	return false
}

func AuthenticateByEmailPassword(email string, password string, userid int) bool {
	req := library.LoginRequest{
		Email:    email,
		Password: password,
	}
	response := tryLogin(req)
	if len(response.Error) == 0 {
		return AuthenticateByToken(response.Token, userid)
	}
	return false
}

func updateQuota(userid int, quota int) string {
	_, err := DB.Exec("update user_information set qouta = ? where userid = ?", quota, userid)
	if err != nil {

		fmt.Println(err)
		return err.Error()
	}
	return ""
}

func addResourceCount(userId int, val int) (error string){
	_, err := DB.Exec("update user_information set resource_count = resource_count + ? where userid = ?", val, userId)
	if err != nil {

		fmt.Println(err)
		return err.Error()
	}
	return ""
}

func canAddResource(userId int) bool {
	row, err := DB.Query("SELECT resource_count, qouta FROM user_information WHERE userid = ? ", userId)
	if err != nil {
		panic(err)
	}

	var resourceCount, qouta int

	for row.Next() {

		err = row.Scan(&resourceCount, &qouta)
		if err != nil {
			panic(err)
		}
	}
	if qouta == -1{
		return true
	}else if qouta>resourceCount{
		return true
	}
	return false


}

func addResource(userid int, resource string) (error string) {
	list, errorText := listResource(userid)
	if len(errorText) > 0 {
		return errorText
	}
	list = append(list, resource)


	error = modifyResource(userid, list)

	if len(error) > 0 {
		return error
	}
	error = addResourceCount(userid,1)
	if len(error) > 0 {
		return error
	}

	return ""
}

func listResource(userId int) (ret []string, error string) {
	row, err := DB.Query("SELECT resource FROM user_information WHERE userid = ?", userId)
	if err != nil {
		panic(err)
	}
	data := make([]byte, 1024)
	for row.Next() {
		err = row.Scan(&data)
		if err != nil {
			panic(err)
		}
	}
	if data != nil{
		err = json.Unmarshal(data, &ret)
		if err != nil {
			panic(err)
		}
	}


	return ret, ""

}

func modifyResource(userID int, list []string) (error string) {

	data, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	_, err = DB.Exec("update user_information set resource = ? where userid = ?", data, userID)
	if err != nil {

		fmt.Println(err)
		return err.Error()
	}
	return ""
}

func deleteResource(userId int, resource string) (error string) {
	list, errorText := listResource(userId)
	if len(errorText) > 0 {
		return errorText
	}

	var newList []string
	numberOfdeletedResource :=0
	for _, s := range list {
		if s == resource {
			numberOfdeletedResource++
			continue
		}
		newList = append(newList, s)
	}

	error = modifyResource(userId, newList)

	if len(error) > 0 {
		return error
	}
	error = addResourceCount(userId,-numberOfdeletedResource)
	if len(error) > 0 {
		return error
	}

	return ""

}

func createUser(email string, password string) (error string) {
	hashedValue := getHashValue(password);
	_, err := DB.Exec("insert into user_information(email, password) values(?, ?)", email, hashedValue)
	if err != nil {

		fmt.Println(err)
		return err.Error()
	}
	return ""

}

func listUser(limit int, offset int) (list []library.User, error string) {
	row, err := DB.Query("SELECT userid, email, usertype, deleted,resource, resource_count, qouta FROM user_information limit ? , ?  ", offset, limit)
	if err != nil {
		panic(err)
	}

	data := make([]byte, 1024)
	for row.Next() {
		cur := library.User{}
		err = row.Scan(&cur.UserID,
			&cur.Email,
			&cur.UserType,
			&cur.Deleted,
			&data,
			&cur.ResourceCount,
			&cur.Quota)
		if err != nil {
			panic(err)
			return list, err.Error()
		}
		if data != nil {
			err = json.Unmarshal(data, &cur.Resource)
			if err != nil {
				panic(err)
				return list, err.Error()
			}
		}
		list = append(list, cur)

	}
	return list, ""
}

func deleteUser(userId int) (error string) {
	_, err := DB.Exec("update user_information set deleted = 1 where userid = ?", userId)
	if err != nil {

		fmt.Println(err)
		return err.Error()
	}
	return ""
}
