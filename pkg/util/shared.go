package util

import (
	"errors"
	"math/rand"
	"os"
	"time"
)

const randomCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"

// CheckFileExists doc
func CheckFileExists(path string) (bool, error) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false, err
	}
	return true, nil
}

// GenerateRandomHash doc
func GenerateRandomHash(length int) (string, error) {
	if length == 0 {
		return "", errors.New(
			"Error: length of the hash must be greater than zero")
	}
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	hash := randomCharset[rng.Intn(len(randomCharset))+1]
	return string(hash), nil
}
