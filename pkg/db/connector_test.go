package db

import (
	"golock3r/server/authtool"
	"golock3r/server/logger"
	"os"
	"strings"
	"testing"
)

func removeFiles() {
	os.Remove("testlogins.txt")
	os.Remove("testlogs.txt")
}

// Helper method to receive a unique key from a validated user
func getKey() []byte {
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.LoginFile = "testlogins.txt"
	authtool.Loggers = Loggers
	authtool.CreateUser("Test", "Password")
	return authtool.GetKey("Test", "Password")
}

func TestConnect(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	if !Connect("demo") {
		t.Error("Failed to connect to database")
	}
}

func TestEncryptDecryptEntry(t *testing.T) {
	entry := map[string]string{
		"title":        "Test Title",
		"password":     "VerySecurePassword",
		"username":     "Test",
		"private_note": "There is a private note here! Don't tell your dad!",
		"public_note":  "This is a public note. Feel free to share with your pop!",
	}

	key := getKey()

	enc_entry := EncryptEntry(key, entry)

	for k := range entry {
		if k == "password" || strings.Contains(k, "private") {
			if enc_entry[k] == entry[k] {
				t.Error("Not encrypting passwords or 'private' keys")
			} else if len([]byte(enc_entry[k]))%16 != 0 {
				t.Error("Hey, this encrypted block is either too long or too short. Get Goldilocks in here.")
			}
		} else {
			if enc_entry[k] != entry[k] {
				t.Error("Hey, this plaintext field is encrypted. What happened?")
			}
		}
	}
}

func TestWriteEntry(t *testing.T) {

	Connect("test")

	entry := map[string]string{
		"title":        "Test Title",
		"password":     "VerySecurePassword",
		"username":     "Test",
		"private_note": "There is a private note here! Don't tell your dad!",
		"public_note":  "This is a public note. Feel free to share with your pop!",
	}
	entry2 := map[string]string{
		"title":        "Title",
		"password":     "VerySecurePassword",
		"username":     "Test1",
		"private_note": "There is a private note here! Don't tell your dad!",
		"public_note":  "This is a public note. Feel free to share with your pop!",
	}

	we := WriteEntry(entry)
	we2 := WriteEntry(entry2)
	if we != true && we2 != true {
		t.Error("Entry was not written entry was exepected to be written ")
	}
}

func TestReadFromTitle(t *testing.T) {
	Connect("test")
	rt := ReadFromTitle("Test Title")

	if rt == nil || rt[0]["title"] != "Test Title" {
		t.Error("expected one result got more than one or no result")
	}
}

func TestReadFromUsername(t *testing.T) {
	Connect("test")
	ru := ReadFromUsername("Test")

	if ru == nil || ru[0]["username"] != "Test" {
		t.Error("unexpected results expected to find entry found either a null entry or the wrong entry")
	}
}

func TestReadAll(t *testing.T) {
	Connect("test")

	ra := ReadAll()

	if ra == nil || len(ra) < 2 {

		t.Error("unexpected results expected to find two")
	}
}

func TestUpdate(t *testing.T) {

	if !UpdateEntry("title", "title", "testUpdate") {
		t.Error("unable to update database entry")
	}

}

func TestDelete(t *testing.T) {
	if !DeleteEntry("Test") {
		t.Error("unable to delete")
	}

}

func TestRemoveAll(t *testing.T) {
	Connect("test")
	if !RemoveAll() {
		t.Error("unable to remove all entries")
	}
	removeFiles()
}
