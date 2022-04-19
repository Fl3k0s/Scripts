package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"
)

//user and password for the database
var users = [...]string{"EXAMPLE1", "EXAMPLE2"}
var passwords = [...]string{"A", "B"}
var x int
var date = exportDate()

//this is a query for couchbase
var query = "INSERT INTO ecom-pro.auth.authentication (KEY, VALUE) VALUES "

func main() {
	x = len(users)

	//queryF := generateQuerysCouchbase(query)
	json := _generateJson()
	sql := _generateSql()
	splitDate()

	sqlFileName := "./files/inesertQuery-" + date + ".sql"
	jsonFileName := "./files/usersCouchbase-" + date + ".json"

	_generateFile(json, jsonFileName)
	_generateFile(sql, sqlFileName)

	fmt.Println(json)
	fmt.Println(sql)
}

//split text in lines
func splitDate() {
	var lines []string
	var textSplit string
	for _, line := range strings.Split(date, "-") {
		line = strings.TrimSpace(line)

		if line != "" {
			lines = append(lines, line)
		}

		textSplit = textSplit + line
	}

	date = textSplit
	fmt.Println(date)
	fmt.Println(textSplit)
}

//time.Time to string
func exportDate() string {
	t := time.Now()
	return t.Format("2006-01-02")
}

//generate a sql file, import the sql text
func _generateFile(text string, fileName string) {
	f, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		return

	}

	l, err := f.WriteString(text)

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
func _generateSql() string {
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

//encode the password
func _encodePassword(password string) string {
	encrypted := sha256.Sum256([]byte(password))
	ps := hex.EncodeToString(encrypted[:])
	return ps
}

//this method format the user and password for the json text
func _formatUserAndPassword(user, password string) (userF string, passF string) {
	userF = "\"" + user + "\""
	passF = "\"" + password + "\""
	return
}

//this method generate json text
//the length of arrays is the value of x
func _generateJson() string {
	json := "[\n"

	for i := 0; i < x; i++ {
		ps := _encodePassword(passwords[i])

		user, pass := _formatUserAndPassword(users[i], ps)

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
		ps := _encodePassword(passwords[i])

		user, pass := _formatUserAndPassword(users[i], ps)

		value := "(" + users[i] + ",{\"username\"" + " : " + user + " , \"password\" : " + pass

		//fmt.Println(value)

		//this is the final query
		queryF := query + value + "});\n"

		// all querys in only text
		querys = querys + queryF
	}

	return querys
}
