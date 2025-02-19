package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomRequestId() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)
}
