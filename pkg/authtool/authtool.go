// Contains tools for creating, deleting, editing and authenticating users
package authtool

import (
	"bufio"
	"crypto/sha512"
	"golock3r/server/logger"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Structure for deleting and modifying entries
type entry struct {
	user string
	pass []byte
}

var LoginFile = "logins.txt"
var Loggers *logger.Loggers

var Rounds = 12

// Given a password and number of bcrypt rounds, hash the given password
// and return its hash value.
func HashUserPassword(password string) []byte {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), Rounds)
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
		file, err := os.OpenFile(File, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0700)

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
		return WriteFile(username, HashUserPassword(password))
	} else {
		return false
	}

}

// Given a valid username and password, change the associated password.
// Given a valid username, delete the user's username and password entry
// from the logins file. Returns true if either is successful, false if otherwise
func ModifyUser(username string, password string, delete_user bool) bool {
	file, err := os.Open(LoginFile)

	if err != nil {
		Loggers.LogError.Println("Problem opening / reading file", err)
		return false
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var entries []entry

	for scanner.Scan() {

		line := scanner.Text()
		logins_user := strings.Split(line, ":")[0]
		logins_pass := []byte(strings.Split(line, ":")[1])

		if delete_user {
			if !(logins_user == username) {
				entries = append(entries, entry{logins_user, logins_pass})
			}
		} else {
			if logins_user == username {
				entries = append(entries, entry{logins_user, HashUserPassword(password)})
			} else {
				entries = append(entries, entry{logins_user, logins_pass})
			}
		}

	}

	err = os.Remove(LoginFile)

	if err != nil {
		Loggers.LogError.Println("Problems deleting user from logins file! Cannot re-create file!")
		return false
	} else {
		for _, ent := range entries {
			WriteFile(ent.user, ent.pass)
		}
	}
	Loggers.LogInfo.Println("Modified / deleted user")
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

// Returns a SHA512 sum used as the AES encryption key for an authenticated user
// Returns nothing if the user is not authenticated
func GetKey(username string, password string) []byte {
	sha_hash := sha512.New()
	var keystring = username + password

	if ValidateUser(username, password) {
		sha_hash.Write([]byte(keystring))
		return sha_hash.Sum(nil)[0:32]
	} else {
		return nil
	}
}
