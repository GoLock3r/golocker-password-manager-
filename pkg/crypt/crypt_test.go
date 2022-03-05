package crypt

import (
	"golock3r/server/authtool"
	"golock3r/server/logger"
	"testing"
)

var pt = "a"

func TestEncryptDecrypt(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers

	authtool.CreateUser("Test", "Password")
	key := authtool.GetKey("Test", "Password")

	chunk_data := ChunkStringData(pt)

	ret := Encrypt(key, chunk_data)
	t.Errorf(CleanStringData(ret))

	ret2 := Decrypt(key, ret)
	t.Errorf(CleanStringData(ret2))
}
