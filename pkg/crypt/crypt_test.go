package crypt

import (
	"golock3r/server/authtool"
	"golock3r/server/logger"
	"os"
	"testing"
)

var sample_1 = "The quick brown fox jumps over the lazy dog."
var sample_2 = "SomeUsernameHere"
var sample_3 = "VerySecurePassword!123"

func removeFiles() {
	os.Remove("testlogins.txt")
	os.Remove("testlogs.txt")
}

// Helper method to test the equality of two byte arrays
func checkEquality(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
	}
	return true
}

// Helper method to receive a unique key from a validated user
func getKey() []byte {
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.LoginFile = "testlogins.txt"
	authtool.Loggers = Loggers
	authtool.CreateUser("Test", "Password")
	return authtool.GetKey("Test", "Password")
}

func TestChunkData(t *testing.T) {
	chunk_sample_1 := ChunkStringData(sample_1)
	chunk_sample_2 := ChunkStringData(sample_2)
	chunk_sample_3 := ChunkStringData(sample_3)
	// Test sample 1
	for _, block := range chunk_sample_1 {
		if len(block) != 16 {
			t.Errorf("Chunks should be of size 16")
		}
	}
	// Test sample 2
	for _, block := range chunk_sample_2 {
		if len(block) != 16 {
			t.Errorf("Chunks should be of size 16")
		}
	}
	// Test sample 3
	for _, block := range chunk_sample_3 {
		if len(block) != 16 {
			t.Errorf("Chunks should be of size 16")
		}
	}
}

func TestStringFunctions(t *testing.T) {
	chunk_sample_1 := ChunkStringData(sample_1)
	chunk_sample_2 := ChunkStringData(sample_2)
	chunk_sample_3 := ChunkStringData(sample_3)
	// Test sample 1
	if CleanStringData(chunk_sample_1) != sample_1 {
		t.Errorf("CleanStringData() not clearing padding or some other issue.")
	}
	// Test sample 2
	if CleanStringData(chunk_sample_2) != sample_2 {
		t.Errorf("CleanStringData() not clearing padding or some other issue.")
	}
	// Test sample 3
	if CleanStringData(chunk_sample_3) != sample_3 {
		t.Errorf("CleanStringData() not clearing padding or some other issue.")
	}
	// Should not equal as ToString does not remove padding
	if ToString(chunk_sample_1) == sample_1 {
		t.Errorf("ToString should not remove padding")
	}
}

func TestFormats(t *testing.T) {
	chunk_sample_1 := ChunkStringData(sample_1)
	storage_fmt := FormatStorage(chunk_sample_1)
	raw_fmt := FormatRaw(storage_fmt)
	// raw_fmt should be the same as chunk_sample_1
	if ToString(raw_fmt) == ToString(chunk_sample_1) {
		t.Errorf("Improper formatting from FormatRaw()")
	}
	// FormatRaw() should return 32 byte blocks
	for _, block := range raw_fmt {
		if len(block) != 32 {
			t.Errorf("FormatRaw() returns blocks of incorrect length. Blocks should be 32 bytes.")
		}
	}
}

func TestEncryptDecrypt(t *testing.T) {
	key := getKey()

	chunk_sample_1 := ChunkStringData(sample_1)
	chunk_sample_2 := ChunkStringData(sample_2)
	chunk_sample_3 := ChunkStringData(sample_3)

	enc1 := Encrypt(key, chunk_sample_1)
	enc2 := Encrypt(key, chunk_sample_2)
	enc3 := Encrypt(key, chunk_sample_3)

	if ToString(enc1) == ToString(chunk_sample_1) {
		t.Errorf("Encrypted output shouldn't be equal to plaintext input")
	}

	if ToString(enc2) == ToString(chunk_sample_2) {
		t.Errorf("Encrypted output shouldn't be equal to plaintext input")
	}

	if ToString(enc3) == ToString(chunk_sample_3) {
		t.Errorf("Encrypted output shouldn't be equal to plaintext input")
	}

	dec1 := Decrypt(key, enc1)
	dec2 := Decrypt(key, enc2)
	dec3 := Decrypt(key, enc3)

	if ToString(dec1) != ToString(chunk_sample_1) {
		t.Errorf("Improper decryption. Decrypted output does not equal input")
	}

	if ToString(dec2) != ToString(chunk_sample_2) {
		t.Errorf("Improper decryption. Decrypted output does not equal input")
	}

	if ToString(dec3) != ToString(chunk_sample_3) {
		t.Errorf("Improper decryption. Decrypted output does not equal input")
	}
}

func TestEncryptDecryptFormatting(t *testing.T) {
	key := getKey()
	chunk_sample_1 := ChunkStringData(sample_1)
	enc1_fmt := FormatStorage(Encrypt(key, chunk_sample_1))
	enc1_raw := FormatRaw(enc1_fmt)

	dec1 := Decrypt(key, enc1_raw)

	if ToString(dec1) != ToString(chunk_sample_1) {
		t.Errorf("Formatting is interfering with decryption. Decrypted output does not equal input")
	}
}

func TestNewChanges(t *testing.T) {
	key := getKey()
	chunk_sample_1 := ChunkStringData(sample_2)

	enc1_fmt := FormatHex(Encrypt(key, chunk_sample_1))
	enc1_raw := FormatHexToRaw(enc1_fmt)
	t.Error(enc1_fmt)
	t.Error(HexToString(enc1_raw))

	dec1 := Decrypt(key, enc1_raw)
	t.Error(CleanStringData(dec1), CleanStringData(chunk_sample_1))

	e := EncryptStringToHex(key, "Testing")
	t.Error(e)
	t.Error(DecryptStringFromHex(key, e))
}

func TestCleanup(t *testing.T) {
	// Remove test logs & login files at the end of execution
	removeFiles()
}
