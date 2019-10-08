package handlers

import (
	"database/sql"
	"encoding/json"
	"time"

	"../utilities"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func GetDB() *sql.DB {
	dbDetails := utilities.Config.DBUsername + ":" +
		utilities.Config.DBPassword + "@tcp(" +
		utilities.Config.DBAddress + ":" +
		utilities.Config.DBPort + ")/" +
		utilities.Config.DBName + "?timeout=1s"
	Db, err := sql.Open("mysql", dbDetails)
	if err != nil {
		Log.WarningF("Error connecting to DB. %s", err.Error())
		return nil
	}
	return Db
}

func tryLogin(request utilities.LoginRequest) (response utilities.LoginResponse) {

	response.Success = false
	row, err := DB.Query("SELECT password, userid FROM user_information WHERE email = ? and deleted = 0", request.Email)
	if err != nil {
		Log.WarningF("Error DB Query for login. %s", err.Error())
		response.Error = err.Error()
		return
	}

	var L string
	var id int

	for row.Next() {

		err = row.Scan(&L, &id)
		if err != nil {
			Log.WarningF("Error Scanning values. %s", err.Error())
			response.Error = err.Error()
		}
	}

	passwordMatched, err := ComparePasswords(L, request.Password)
	if err != nil {
		response.Error = "Mismatch Password"
		return
	}
	if passwordMatched {
		token, err := getToken(request.Email)
		if err != nil {
			response.Error = "Token Error"
			Log.WarningF("Error getting Token. %s", err.Error())
			return
		}
		response.UserID = id
		response.Token = token
		response.Success = true
		return
	} else {
		Log.InfoF("Authentication Error for Email: %s", request.Email)
		response.Error = "Wrong email/password"
		return
	}

}

func getToken(email string) (token string, err error) {

	token = IssueToken(email, 24*time.Hour)
	_, err = DB.Exec("update user_information set token = ? where email = ?", token, email)
	if err != nil {
		Log.WarningF("Error getting Token. %s", err.Error())
		return "", err
	}
	return token, nil
}

func AuthenticateByToken(token string, userid int) bool {

	ok, email := VerifyToken(token)
	if !ok {
		Log.InfoF("Token Verification False for email: %s", email)
		return false
	}

	row, err := DB.Query("SELECT userid, usertype FROM user_information WHERE email = ? and deleted = 0", email)
	if err != nil {
		Log.WarningF("Error DB Query for Token Authentication. %s", err.Error())
	}
	var id, userType int

	for row.Next() {

		err = row.Scan(&id, &userType)
		if err != nil {
			Log.WarningF("Error scanning values. %s", err.Error())
			return false
		}
	}

	if userType == userAdmin || id == userid {
		return true
	}
	return false
}

func AuthenticateByEmailPassword(email string, password string, userid int) bool {
	req := utilities.LoginRequest{
		Email:    email,
		Password: password,
	}
	response := tryLogin(req)

	if len(response.Error) == 0 {
		return AuthenticateByToken(response.Token, userid)
	}
	return false
}

func updateQuota(userid int, quota int) (err error) {

	_, err = DB.Exec("update user_information set qouta = ? where userid = ?", quota, userid)
	if err != nil {
		Log.WarningF("Error updating Quota in DB. %s", err.Error())
		return err
	}
	return nil
}

func addResourceCount(userId int, val int) (err error) {
	_, err = DB.Exec("update user_information set resource_count = resource_count + ? where userid = ?", val, userId)
	if err != nil {
		Log.WarningF("Error adding Resource in DB. %s", err.Error())
		return err
	}
	return nil
}

func canAddResource(userId int) bool {
	row, err := DB.Query("SELECT resource_count, qouta FROM user_information WHERE userid = ? ", userId)
	if err != nil {
		Log.WarningF("Error DB Query for Resource. %s", err.Error())
		return false
	}

	var resourceCount, qouta int

	for row.Next() {

		err = row.Scan(&resourceCount, &qouta)
		if err != nil {
			Log.WarningF("Error scanning for values. %s", err.Error())
		}
	}

	if qouta == -1 {
		return true
	} else if qouta > resourceCount {
		return true
	}
	return false

}

func addResource(userid int, resource string) (err error) {

	list, err := listResource(userid)
	if err != nil {
		return err
	}
	list = append(list, resource)

	err = modifyResource(userid, list)

	if err != nil {
		return err
	}

	err = addResourceCount(userid, 1)
	if err != nil {
		return err
	}
	return nil
}

func listResource(userId int) (ret []string, err error) {
	row, err := DB.Query("SELECT resource FROM user_information WHERE userid = ?", userId)
	if err != nil {
		Log.WarningF("Error DB Query for list Resource. %s", err.Error())
		return
	}
	data := make([]byte, 1024)
	for row.Next() {
		err = row.Scan(&data)
		if err != nil {
			Log.WarningF("Error Scanning values. %s", err.Error())
		}
	}
	if data != nil {
		err = json.Unmarshal(data, &ret)
		if err != nil {
			Log.WarningF("Error Json Unmarshal. %s", err.Error())
		}
	}

	return

}

func modifyResource(userID int, list []string) (err error) {

	data, err := json.Marshal(list)
	if err != nil {
		Log.WarningF("Error Json Unmarshal. %s", err.Error())
		return err
	}

	_, err = DB.Exec("update user_information set resource = ? where userid = ?", data, userID)
	if err != nil {

		Log.WarningF("Error DB Execution. %s", err.Error())
		return err
	}
	return nil
}

func deleteResource(userId int, resource string) (err error) {
	list, err := listResource(userId)
	if err != nil {
		return
	}

	var newList []string
	numberOfdeletedResource := 0
	for _, s := range list {
		if s == resource {
			numberOfdeletedResource++
			continue
		}
		newList = append(newList, s)
	}

	err = modifyResource(userId, newList)
	if err != nil {
		return
	}

	err = addResourceCount(userId, -numberOfdeletedResource)
	if err != nil {
		return
	}
	return nil
}

func createUser(email string, password string, userType int) (err error) {

	hashedPassword, err := HashAndSalt(password)
	if err != nil {
		return
	}
	_, err = DB.Exec("insert into user_information(email, password, usertype) values(?, ?, ?)", email, hashedPassword, userType)
	if err != nil {
		Log.WarningF("Error inserting into DB. %s", err.Error())
		return err
	}
	return nil

}

func listUser(limit int, offset int) (list []utilities.User, err error) {
	row, err := DB.Query("SELECT userid, email, usertype, deleted,resource, resource_count, qouta FROM user_information limit ? , ?  ", offset, limit)
	if err != nil {
		Log.WarningF("Error DB Query. %s", err.Error())
		return
	}

	data := make([]byte, 1024)
	for row.Next() {
		cur := utilities.User{}
		err = row.Scan(&cur.UserID,
			&cur.Email,
			&cur.UserType,
			&cur.Deleted,
			&data,
			&cur.ResourceCount,
			&cur.Quota)
		if err != nil {
			Log.WarningF("Error Scanning Values. %s", err.Error())
			return
		}
		if data != nil {
			err = json.Unmarshal(data, &cur.Resource)
			if err != nil {
				Log.WarningF("Error Json Unmarshal. %s", err.Error())
				return
			}
		}
		list = append(list, cur)

	}
	return
}

func deleteUser(userId int) (err error) {
	_, err = DB.Exec("update user_information set deleted = 1 where userid = ?", userId)
	if err != nil {

		Log.WarningF("Error DB Execution. %s", err.Error())
		return err
	}
	return nil
}
