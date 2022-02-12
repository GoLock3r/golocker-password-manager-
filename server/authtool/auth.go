package authtool

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

/*
Basic flow:
	Authenticator -> Database interface -> access entries (read, add, edit) -> deliver to user
*/

const LOG_FILE = "logs.txt"
const PASSWD_FILE = "logins.txt"

var hash []byte
var user string

var (
	WarnLog  *log.Logger
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func Init() {
	// https://www.honeybadger.io/blog/golang-logging/
	file, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	WarnLog = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLog = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	hash = []byte("0")
	user = ""
}

func HashUserPassword(username string, password string, rounds int) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), rounds)
	if err != nil {
		ErrorLog.Println(err)
	} else {
		InfoLog.Println("Password hashed successfully")
	}
	hash = bytes
	user = username
	return string(bytes)
}

//https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
func ValidateUser(username string, password string) bool {
	file, err := os.Open(PASSWD_FILE)

	if err != nil {
		ErrorLog.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Split(scanner.Text(), ":")[0] == username {
			err := bcrypt.CompareHashAndPassword([]byte(strings.Split(scanner.Text(), ":")[1]), []byte(password))
			if err == nil {
				return true
			}
		}
	}
	return false
}

//https://varunpant.com/posts/reading-and-writing-binary-files-in-go-lang/

func WriteFile() {
	if string(hash[0]) == "0" {
		WarnLog.Println("")
	} else {
		line := user + ":" + string(hash) + "\n"
		file, err := os.OpenFile(PASSWD_FILE, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

		if err != nil {
			ErrorLog.Println("password file not found / cannot be opened!")
		}

		defer file.Close()

		len, err := file.WriteString(line)

		if err != nil {
			ErrorLog.Println(err)
		} else {
			InfoLog.Println("Wrote " + strconv.Itoa(len) + " bytes of data")
		}
	}
}

// func main() {
// 	var user string

// 	fmt.Print("Enter your username: ")
// 	fmt.Scanln(&user)

// 	fmt.Print("Enter your password: ")
// 	password, _ := terminal.ReadPassword(0)

// 	// hashUserPassword(user, string(password), 20)
// 	// writeFile()

// 	if validateUser(user, string(password)) {
// 		fmt.Println("\nYou're in!")
// 		InfoLog.Println("A user has logged in")
// 	} else {
// 		fmt.Println("\nInvalid username / password")
// 		WarnLog.Println("Invalid login attempt!")
// 	}
// }
