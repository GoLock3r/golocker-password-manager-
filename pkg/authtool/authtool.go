// Contains tools for creating, deleting, editing and authenticating users
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

// Given a password and number of bcrypt rounds, hash the given password
// and return its hash value.
func HashUserPassword(password string, rounds int) []byte {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), rounds)
	if err != nil {
		Loggers.LogError.Println(err)
	} else {
		Loggers.LogInfo.Println("Password hashed successfully")
	}

	return bytes
}

// Given a username and password, read login file and verify that the user's
// username exists and the password hashes match. Returns true if the user is valid,
// false if otherwise
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

// Write to the login file with the user's username and bcrypt password hash.
// Creates a new file if the file doesn't already exist. Returns true if the
// parameters were successfully written and false if otherwise.
func WriteFile(user string, hash []byte) bool {
	var status = false
	if string(hash[0]) == "0" {
		Loggers.LogWarning.Println("Password hash is empty. Not writing file")
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

// Given a username and password, hash their password and add the user to the
// logins file if the username doesn't already exist in the file. Returns true
// if successful and false if otherwise
func CreateUser(username string, password string) bool {
	if !userExists(username) {
		return WriteFile(username, HashUserPassword(password, 12))
	} else {
		return false
	}

}

// Given a username, delete a user from the logins file. Returns true if the user
// is successfully deleted and false if otherwise
func DeleteUser(username string) bool {
	return true
}

// Given a username and a new password, delete the old password hash and replace it
// with the hash of the new password. Returns true if the user's password is successfully
// changed, false if otherwise
func ChangePassword(username string, new_password string) bool {
	return true
}

// Private helper method which reads the logins file to verify the existence of a user
// in the logins file. Returns true if the user exist and false if otherwise
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
