package crypt

//https://go.dev/src/crypto/cipher/example_test.go

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"golock3r/server/logger"
	"io"
	"strconv"
	"strings"
)

var Loggers *logger.Loggers

// Chunks data into byte arrays of 16 bytes each. Pads 0's
// on remaining chunks and returns an array of chunks
func ChunkStringData(data string) [][]byte {
	raw := []byte(data)
	var block []byte
	var chunks [][]byte
	i := 1

	for _, b := range raw {
		block = append(block, b)

		if i%16 == 0 {
			chunks = append(chunks, block)
			block = nil
		}
		i += 1
	}
	// Pad 0 values to the end of the last block to complete the block
	rem := 16 - len(block)
	if rem != 16 {
		for j := 0; j < rem; j++ {
			block = append(block, 0)
		}
		// Append the remaining chunks
		chunks = append(chunks, block)
	}

	return chunks
}

// Converts raw bytes into a clean string, stripping all padding
func CleanStringData(data [][]byte) string {
	str_encoded := ""
	for _, block := range data {
		for _, b := range block {
			if !(b == byte(0)) {
				str_encoded += string(b)
			}
		}
	}
	return str_encoded
}

// Convert byte data to a string. Does not strip padding
func ToString(data [][]byte) string {
	str := ""

	for _, block := range data {
		for _, b := range block {
			str += string(b)
		}
	}
	return str
}

// Converts byte data into a string of int values. Used for storage
func FormatStorage(data [][]byte) string {
	str := ""

	for _, block := range data {
		for _, b := range block {
			str += strconv.Itoa(int(b)) + " "
		}
	}
	return str
}

// Converts storage data into byte array of 32 byte blocks
func FormatRaw(data string) [][]byte {
	var block []byte
	var chunks [][]byte
	i := 1
	raw := strings.Split(data, " ")
	raw = raw[:len(raw)-1]

	for _, str_val := range raw {
		val, _ := strconv.Atoi(str_val)
		block = append(block, byte(val))
		if i%32 == 0 {
			chunks = append(chunks, block)
			block = nil
		}
		i += 1
	}
	return chunks
}

// Given a key and a byte array of chunked input of plaintext data (16 bytes each),
// encrypt each chunk and return an array of encrypted chunked data (16 bytes each)
func Encrypt(key []byte, data [][]byte) [][]byte {
	var ciphertext [][]byte

	for _, pt := range data {
		block, err := aes.NewCipher(key)

		if err != nil {
			Loggers.LogError.Println("AES Cipher error!", err)
		}

		cp := make([]byte, aes.BlockSize+len(pt))
		iv := cp[:aes.BlockSize]

		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			Loggers.LogError.Println(err)
		}
		mode := cipher.NewCBCEncrypter(block, iv)
		mode.CryptBlocks(cp[aes.BlockSize:], pt)

		ciphertext = append(ciphertext, cp)
	}
	return ciphertext
}

// Given a key and a byte array of chunked encrypted data (16 bytes each),
// decrypt each chunk and return an array of decrypted chunked data (16 bytes each)
func Decrypt(key []byte, data [][]byte) [][]byte {
	var plaintext [][]byte
	for _, ct := range data {
		block, err := aes.NewCipher(key)

		if err != nil {
			Loggers.LogError.Println("AES Cipher error!", err)
		}

		iv := ct[:aes.BlockSize]
		ct := ct[aes.BlockSize:]

		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(ct, ct)
		plaintext = append(plaintext, ct)
	}
	return plaintext
}
