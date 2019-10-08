package handlers

import (
	"../utilities"
	"database/sql"
	"encoding/json"

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

func getPasswordAndUserIDFromDB(email string)(password string, userId int, err error){

	row, err := DB.Query("SELECT password, userid FROM user_information WHERE email = ? and deleted = 0", email)
	if err != nil {
		Log.WarningF("Error DB Query for login. %s", err.Error())
		return
	}
	for row.Next() {
		err = row.Scan(&password, &userId)
		if err != nil {
			Log.WarningF("Error Scanning values. %s", err.Error())
			return
		}
	}
	return
}


func updateToken(token string, email string) (err error) {

	_, err = DB.Exec("update user_information set token = ? where email = ?", token, email)
	if err != nil {
		Log.WarningF("Error getting Token. %s", err.Error())
		return err
	}
	return nil
}

func getUserIdAndUserTypeFromDB(email string)(userId int, userType int, err error){

	row, err := DB.Query("SELECT userid, usertype FROM user_information WHERE email = ? and deleted = 0", email)
	if err != nil {
		Log.WarningF("Error DB Query for Token Authentication. %s", err.Error())
		return
	}

	for row.Next() {

		err = row.Scan(&userId, &userType)
		if err != nil {
			Log.WarningF("Error scanning values. %s", err.Error())
			return
		}
	}
	return
}


func updateQuota(userId int, quota int) (err error) {

	_, err = DB.Exec("update user_information set qouta = ? where userid = ?", quota, userId)
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

func addResource(userId int, resource string) (err error) {

	list, err := listResource(userId)
	if err != nil {
		return err
	}
	list = append(list, resource)

	err = modifyResource(userId, list)

	if err != nil {
		return err
	}

	err = addResourceCount(userId, 1)
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
	numberOfDeletedResource := 0
	for _, s := range list {
		if s == resource {
			numberOfDeletedResource++
			continue
		}
		newList = append(newList, s)
	}

	err = modifyResource(userId, newList)
	if err != nil {
		return
	}

	err = addResourceCount(userId, -numberOfDeletedResource)
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
