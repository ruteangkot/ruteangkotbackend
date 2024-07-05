package helper

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateResetToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
