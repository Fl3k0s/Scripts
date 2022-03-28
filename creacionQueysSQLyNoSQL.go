package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

//user and password for the database
var users = [...]string{"A", "B", "C"}
var passwords = [...]string{"A", "B", "C"}
var x int

//this is a query for couchbase
var query = "INSERT INTO ecom-pro.auth.authentication (KEY, VALUE) VALUES "

func main() {
	//queryF := generateQuerysCouchbase(query)

	x = len(users)
	json := generateJson()
	sql := generateSql()

	generateJsonFile(json)
	generateSqlFile(sql)

	fmt.Println(json)
	fmt.Println(sql)
}

//generate a sql file, import the sql text
func generateSqlFile(sql string) {
	f, err := os.Create("./files/query.sql")

	if err != nil {
		fmt.Println(err)
		return

	}

	l, err := f.WriteString(sql)

	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")

	err = f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
}

//this method generate sql query
//the length of array is the value of x
func generateSql() string {
	query := "INSERT INTO shipping.DRIVER_GEOLOCATION (DRIVER_ID, LATITUDE, LONGITUDE) \nVALUES "

	for i := 0; i < x; i++ {
		value := "('" + users[i] + "', 0.0, 0.0)"

		if i != x-1 {
			value = value + ",\n"
		}

		query = query + value
	}

	query = query + ";"

	return query
}

//generate a json file
func generateJsonFile(json string) {
	f, err := os.Create("./files/users.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString(json)

	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")

	err = f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
}

//encode the password
func encodePassword(password string) string {
	encrypted := sha256.Sum256([]byte(password))
	ps := hex.EncodeToString(encrypted[:])
	return ps
}

//this method format the user and password for the json text
func formatUserAndPassword(user, password string) (userF string, passF string) {
	userF = "\"" + user + "\""
	passF = "\"" + password + "\""
	return
}

//this method generate json text
//the length of arrays is the value of x
func generateJson() string {
	json := "[\n"

	for i := 0; i < x; i++ {
		ps := encodePassword(passwords[i])

		user, pass := formatUserAndPassword(users[i], ps)

		value := "\t{\n\t\t\"username\"" + " : " + user + " ,\n\t\t\"password\" : " + pass + "\n\t}"

		if i != x-1 {
			value = value + ",\n"
		}

		json = json + value
	}

	json = json + "\n]"
	return json
}

//normaly don't use this
//the length of arrays is the value of x
func generateQuerysCouchbase() string {
	querys := ""

	for i := 0; i < x; i++ {
		ps := encodePassword(passwords[i])

		user, pass := formatUserAndPassword(users[i], ps)

		value := "(" + users[i] + ",{\"username\"" + " : " + user + " , \"password\" : " + pass

		//fmt.Println(value)

		//this is the final query
		queryF := query + value + "});\n"

		// all querys in only text
		querys = querys + queryF
	}

	return querys
}
