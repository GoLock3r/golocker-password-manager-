package db

import (
	"golock3r/server/authtool"
	"golock3r/server/logger"
	"testing"
)

func Test(t *testing.T) {
	authtool.Loggers = logger.CreateLoggers("testlogs.txt")
	Loggers = authtool.Loggers

	entry := map[string]string{
		"title":        "Test Title",
		"password":     "VerySecurePassword",
		"username":     "Test",
		"private_note": "There is a private note here!",
		"public_note":  "This is a public note. This will not be encrypted!",
	}

	authtool.CreateUser("test", "test")
	key := authtool.GetKey("test", "test")
	enc_entry := EncryptEntry(key, entry)
	t.Error(enc_entry)
	t.Error(DecryptEntry(key, enc_entry))
	t.Error(DecryptEntry([]byte("11111111111111111111111111111111"), enc_entry))

	Connect("test")
	WriteEntry(entry)
	WriteEntry(enc_entry)
}
