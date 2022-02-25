// Test authtool package memebers
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
	// Verify hash is returned
	result := HashUserPassword(password, 12)
	if !strings.Contains(string(result), "$2a$12$") {
		t.Errorf("Bcrypt hash identifier not found. Got %q, wanted %q", result, "$2a$12$*")
	}
}

func TestUserValidation(t *testing.T) {
	removeFiles()

	Loggers = logger.CreateLoggers("testlogs.txt")
	LoginFile = "testlogins.txt"
	// Attempt to create a user
	if !CreateUser(username, password) {
		t.Errorf("Could't create user")
	}
	// Verify that validate user works as intended
	if !ValidateUser(username, password) {
		t.Errorf("ValidateUser() false negative. Got false, wanted true")
	}
	// Check false positives
	if ValidateUser(username, "InvalidPassword") {
		t.Errorf("ValidateUser() false positive on invalid password. Got true, wanted false")
	}
	// Check false positives
	if ValidateUser("InvalidUsername", password) {
		t.Errorf("ValidateUser() false positive on invalid username. Got true, wanted false")
	}
}

func TestDuplicateUserCreation(t *testing.T) {
	removeFiles()

	Loggers = logger.CreateLoggers("testlogs.txt")
	LoginFile = "testlogins.txt"
	// Create a user
	if !CreateUser(username, password) {
		t.Errorf("Could't create user")
	}
	// Create the same user again (should not happen)
	if CreateUser(username, password) {
		t.Errorf("")
	}
}

func TestDeleteUser(t *testing.T) {
	removeFiles()

	Loggers = logger.CreateLoggers("testlogs.txt")
	// Create a user
	if !CreateUser(username, password) {
		t.Errorf("Could't create user")
	}
	// Delete user
	if !DeleteUser(username) {
		t.Errorf("Couldn't delete user")
	}

}

func TestChangePassword(t *testing.T) {
	removeFiles()

	Loggers = logger.CreateLoggers("testlogs.txt")
	// Create a user
	if !CreateUser(username, password) {
		t.Errorf("Could't create user")
	}
	// Change user's password
	if !ChangePassword(username, "NewPassword") {
		t.Errorf("Couldn't update user " + username)
	}
	// Validate user with new password
	if ValidateUser(username, "NewPassword") {
		t.Errorf("Couldn't validate user with updated password")
	}
}

func TestCleanup(t *testing.T) {
	// Remove test logs & login files at the end of execution
	removeFiles()
}
