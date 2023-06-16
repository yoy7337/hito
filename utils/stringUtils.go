package utils

import (
	"crypto/rand"
	"encoding/hex"
	"math"
)

// generate random string with length n
func GenerateRandomHexString(n int) (string, error) {
	b, err := GenerateRandomBytes(int(math.Ceil(float64(n) / 2)))

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

// generate random bytes with length n
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}
