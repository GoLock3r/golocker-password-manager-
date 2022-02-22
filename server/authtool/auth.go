package authtool

import (
	"bufio"
	"golock3r/server/logger"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var LoginFile = "logins.txt"

var hash []byte
var user string

var Loggers *logger.Loggers

func HashUserPassword(username string, password string, rounds int) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), rounds)
	if err != nil {
		Loggers.LogError.Println(err)
	} else {
		Loggers.LogInfo.Println("Password hashed successfully")
	}
	hash = bytes
	user = username
	return string(bytes)
}

//https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
func ValidateUser(username string, password string) bool {
	file, err := os.Open(LoginFile)

	if err != nil {
		Loggers.LogError.Println(err)
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
		Loggers.LogWarning.Println("")
	} else {
		line := user + ":" + string(hash) + "\n"
		file, err := os.OpenFile(LoginFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

		if err != nil {
			Loggers.LogError.Println("password file not found / cannot be opened!")
		}

		defer file.Close()

		len, err := file.WriteString(line)

		if err != nil {
			Loggers.LogError.Println(err)
		} else {
			Loggers.LogInfo.Println("Wrote " + strconv.Itoa(len) + " bytes of data")
		}
	}
}
