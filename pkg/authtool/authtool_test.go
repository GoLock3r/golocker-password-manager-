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
	result := HashUserPassword(password)
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
	LoginFile = "testlogins.txt"
	// Create a user
	if !CreateUser("Username_1", "password1") {
		t.Errorf("Could't create user")
	}
	// Create another user
	if !CreateUser("Username_2", "password2") {
		t.Errorf("Could't create user")
	}
	// Delete a user
	if !ModifyUser("Username_2", "", true) {
		t.Errorf("Couldn't delete user")
	}
	// Attempt to authenticate as the deleted user (should return false)
	if ValidateUser("Username_2", "password2") {
		t.Errorf("Shouldn't be able to login as a deleted user!")
	}
	// Attempt to authenticate other user to ensure it remains in the logins file (should return true)
	if !ValidateUser("Username_1", "password1") {
		t.Errorf("Cannot login as a valid user")
	}
}

func TestChangePassword(t *testing.T) {
	removeFiles()

	Loggers = logger.CreateLoggers("testlogs.txt")
	LoginFile = "testlogins.txt"
	// Create a user
	if !CreateUser("Username_1", "password1") {
		t.Errorf("Could't create user")
	}
	// Create another user
	if !CreateUser("Username_2", "password2") {
		t.Errorf("Could't create user")
	}
	// Change the first users password
	if !ModifyUser("Username_1", "new_password!", false) {
		t.Errorf("Couldn't change user's password")
	}
	// Attempt to authenticate with the users old password
	if ValidateUser("Username_1", "password1") {
		t.Errorf("Shouldn't be able to login with old password")
	}
	// Attempt to authenticate with the users new password
	if !ValidateUser("Username_1", "new_password!") {
		t.Errorf("Can't log in with new password")
	}
}

func TestCleanup(t *testing.T) {
	// Remove test logs & login files at the end of execution
	removeFiles()
}
