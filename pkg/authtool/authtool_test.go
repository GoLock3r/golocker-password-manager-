package authtool

import (
	"golock3r/server/logger"
	"os"
	"strings"
	"testing"
)

var username = "TestUsername"
var password = "SomeComplexPassword"

func removeFiles() {
	os.Remove("testlogins.txt")
	os.Remove("testlogs.txt")
}

func TestHashUserPasswordBcryptOutput(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")

	result := HashUserPassword(username, password, 12)
	if !strings.Contains(string(result), "$2a$12$") {
		t.Errorf("Bcrypt hash identifier not found. Got %q, wanted %q", result, "$2a$12$*")
	}
}

func TestUserValidation(t *testing.T) {
	removeFiles()

	Loggers = logger.CreateLoggers("testlogs.txt")
	LoginFile = "testlogins.txt"

	if !CreateUser(username, password) {
		t.Errorf("Error in WriteFile()")
	}

	if !ValidateUser(username, password) {
		t.Errorf("ValidateUser() false negative. Got false, wanted true")
	}

	if ValidateUser(username, "InvalidPassword") {
		t.Errorf("ValidateUser() false positive on invalid password. Got true, wanted false")
	}

	if ValidateUser("InvalidUsername", password) {
		t.Errorf("ValidateUser() false positive on invalid username. Got true, wanted false")
	}
}

func TestDuplicateUserCreation(t *testing.T) {
	removeFiles()

	Loggers = logger.CreateLoggers("testlogs.txt")
	LoginFile = "testlogins.txt"

	if !CreateUser(username, password) {
		t.Errorf("Error in WriteFile()")
	}

}

func TestCleanup(t *testing.T) {
	removeFiles()
}
