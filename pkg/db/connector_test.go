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

}

func TestReadFromTitle(t *testing.T) {

}

func TestReadFromUsername(t *testing.T) {

}

func TestReadAll(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}

func TestRemoveAll(t *testing.T) {
	removeFiles()
}

// func Test(t *testing.T) {
// 	authtool.Loggers = logger.CreateLoggers("testlogs.txt")
// 	Loggers = authtool.Loggers

// 	entry := map[string]string{
// 		"title":        "Test Title",
// 		"password":     "VerySecurePassword",
// 		"username":     "Test",
// 		"private_note": "There is a private note here! Don't tell your dad!",
// 		"public_note":  "This is a public note. This will not be encrypted!",
// 	}

// 	authtool.CreateUser("test", "test")
// 	key := authtool.GetKey("test", "test")
// 	enc_entry := EncryptEntry(key, entry)
// 	t.Error(enc_entry)
// 	t.Error(DecryptEntry(key, enc_entry))
// 	t.Error(DecryptEntry([]byte("11111111111111111111111111111111"), enc_entry))

// 	Connect("test")
// 	WriteEntry(entry)
// 	WriteEntry(enc_entry)
// }
