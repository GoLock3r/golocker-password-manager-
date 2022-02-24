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

var Loggers *logger.Loggers

func HashUserPassword(username string, password string, rounds int) []byte {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), rounds)
	if err != nil {
		Loggers.LogError.Println(err)
	} else {
		Loggers.LogInfo.Println("Password hashed successfully")
	}

	return bytes
}

//https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
func ValidateUser(username string, password string) bool {
	file, err := os.Open(LoginFile)

	if err != nil {
		Loggers.LogError.Println("Problem opening / reading file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Split(scanner.Text(), ":")[0] == username {
			err := bcrypt.CompareHashAndPassword([]byte(strings.Split(scanner.Text(), ":")[1]), []byte(password))
			if err == nil {
				Loggers.LogInfo.Println("Successful authentication")
				return true
			}
		}
	}
	Loggers.LogInfo.Println("Invalid username / password")
	return false
}

//https://varunpant.com/posts/reading-and-writing-binary-files-in-go-lang/
func WriteFile(user string, hash []byte) bool {
	var status = false
	if string(hash[0]) == "0" {
		Loggers.LogWarning.Println("")
	} else {
		line := user + ":" + string(hash) + "\n"
		file, err := os.OpenFile(LoginFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0700)

		if err != nil {
			Loggers.LogError.Println("Password file not found / cannot be opened.", err)
		}

		defer file.Close()

		len, err := file.WriteString(line)

		if err != nil {
			Loggers.LogError.Println("Opened file but could not write.", err)
		} else {
			Loggers.LogInfo.Println("Wrote " + strconv.Itoa(len) + " bytes of data")
			status = true
		}
	}
	return status
}

func CreateUser(username string, password string) bool {
	if !userExists(username) {
		return WriteFile(username, HashUserPassword(username, password, 12))
	} else {
		return false
	}

}

func DeleteUser(username string) (status bool, response string) {
	return false, "some response here"
}

func userExists(username string) bool {
	file, err := os.Open(LoginFile)

	if err != nil {
		Loggers.LogError.Println("Problem opening / reading file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Split(scanner.Text(), ":")[0] == username {
			return true
		}
	}
	return false
}
