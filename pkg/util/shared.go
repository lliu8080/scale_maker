package util

import (
	"encoding/hex"
	"errors"
	"hash/fnv"
	"os"
	"time"
)

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
	h := fnv.New64a()
	h.Write([]byte(time.Now().Round(time.Hour).String()))
	return hex.EncodeToString(h.Sum(nil))[0:length], nil
}
