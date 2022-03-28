package main
 
import (
	"os"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)
 
func main() {
	users := [...]string{"", "", "", "", "", ""}
	passwords := [...]string{"", "", "", "", "", ""}
	x := 6

	//query := "INSERT INTO ecom-pro.auth.authentication (KEY, VALUE) VALUES "
	//query = printVariousQuerys(users, passwords, x, query)

	json := generateJson(users, passwords, x)
	generateJsonFile(json)
	fmt.Println(json)
}

func generateJsonFile(json string){
	f, err := os.Create("users.json")
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

func encodePassword(password string) string {
	encrypted := sha256.Sum256([]byte(password))
	ps := hex.EncodeToString(encrypted[:])
	return ps
}

func formatUserAndPassword(user, password string)(userF string, passF string) {
	userF = "\"" + user + "\""
	passF = "\"" + password + "\""
	return
}

func generateJson(users [6]string, passwords [6]string, x int) string{
	json := "[\n"
	
	for i := 0 ; i < x ; i++ {
		ps := encodePassword(passwords[i])
		var user string
		var pass string
		user, pass = formatUserAndPassword(users[i], ps)
		value := "\t{\n\t\t\"username\""+ " : " + user + " ,\n\t\t\"password\" : " + pass + "\n\t}"
		if i != x - 1 {
			value = value + ",\n"
		}

		json = json + value
	}

	json = json + "\n]"
	return json
}

func printVariousQuerys(users [6]string, passwords [6]string, x int, query string) string {
	querys := ""
	for i := 0 ; i < x ; i++ {
		ps := encodePassword(passwords[i])
		var user string
		var pass string
		user, pass = formatUserAndPassword(users[i], ps)
		value := "(" + users[i] +",{\"username\""+ " : " + user + " , \"password\" : " + pass
		//fmt.Println(value)
		queryF := query + value + "});\n"
		querys = querys + queryF
	}
	return querys
}