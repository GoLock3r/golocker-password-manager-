package crypt

import (
	"golock3r/server/authtool"
	"golock3r/server/logger"
	"os"
	"testing"
)

var samples = [6]string{"The quick brown fox jumps over the lazy dog.",
	"SomeUsernameHere",
	"VerySecurePassword!123",
	" ",
	"",
	string(byte(0))}

var sample_1 = "The quick brown fox jumps over the lazy dog."
var sample_2 = "SomeUsernameHere"
var sample_3 = "VerySecurePassword!123"

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

func TestChunkDataSize(t *testing.T) {
	// Chunks should be of size 16
	for _, sample := range samples {
		chunk_sample := ChunkStringData(sample)
		for _, block := range chunk_sample {
			if len(block) != 16 {
				t.Errorf("Chunks should be of size 16")
			}
		}

	}
}

func TestChunkDataSizeHex(t *testing.T) {
	// Chunks should be of size 32 when returned from FormatHexToRaw
	for _, sample := range samples {
		chunk_sample := ChunkStringData(sample)
		rawhex_sample := FormatHexToRaw(FormatHex(chunk_sample))
		for _, block := range rawhex_sample {
			if len(block) != 32 {
				t.Errorf("Chunks should be of size 32")
			}
		}
	}
}

func TestStringFunctions(t *testing.T) {
	// Verify that string functions are working as intended
	for _, sample := range samples {
		chunk_sample := ChunkStringData(sample)

		if len(CleanStringData(chunk_sample)) > 0 {
			if len(CleanStringData(chunk_sample)) != len(sample) {
				t.Error("CleanStringData() not clearing padding or some other issue.")
			}
		}

		length_toString := len(ToString(chunk_sample))
		length_sample := len(sample)

		if length_toString != length_sample {
			if ToString(chunk_sample)[length_toString-1] != 0 {
				t.Errorf("ToString should not be removing padding on a block that is not divisible by 16")
			}
		}

	}
}

func TestEncryptDecrypt(t *testing.T) {
	// Test encryption and decryption functionality
	key := getKey()

	for _, sample := range samples {
		chunk_sample := ChunkStringData(sample)

		encrypted_output := Encrypt(key, chunk_sample)

		if len(chunk_sample) != len(encrypted_output) {
			t.Error("Encrypted output does not have the same number of blocks as the chunked sample data")
		} else {
			for i := 0; i < len(chunk_sample); i++ {
				if len(chunk_sample[i]) != len(encrypted_output[i]) {
					if string(chunk_sample[i]) == string(encrypted_output[i]) {
						t.Error("Encrypted output should not be equal to plaintext input")
					}
				} else {
					t.Error("Encrypted output should have 32 byte blocks")
				}
			}
		}

		decrypted_output := Decrypt(key, encrypted_output)

		if len(decrypted_output) != len(chunk_sample) {
			t.Error("Decrypted output does not have the same number of blocks as the chunked sample data")
		} else {
			for i := 0; i < len(chunk_sample); i++ {
				if len(chunk_sample[i]) != len(decrypted_output[i]) {
					t.Error("Decrypted output should have 16 byte blocks")
				} else {
					for j := 0; j < len(chunk_sample[i]); j++ {
						if chunk_sample[i][j] != decrypted_output[i][j] {
							t.Error("Bytes shoud equal! Decryption not returning plaintext", chunk_sample[i][j], decrypted_output[i][j])
						}
					}
				}
			}
		}
	}
}

func TestEncryptDecryptStringToHex(t *testing.T) {
	key := getKey()

	for _, sample := range samples {
		chunk_sample := ChunkStringData(sample)
		enc_hex_str1 := FormatHex(Encrypt(key, chunk_sample))
		enc_hex_str2 := EncryptStringToHex(key, sample)

		enc_chunk_sample := FormatHexToRaw(enc_hex_str1)
		dec_str1 := CleanStringData(Decrypt(key, enc_chunk_sample))
		dec_str2 := DecryptStringFromHex(key, enc_hex_str2)

		if dec_str1 != dec_str2 {
			t.Error("Check EncryptStringToHex & DecryptStringToHex functionality")
		}
	}

}

func TestCleanup(t *testing.T) {
	// Remove test logs & login files at the end of execution
	removeFiles()
}
