package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashText(password string) string {
	hashedBytes := sha256.Sum256([]byte(password))
	hashedString := hex.EncodeToString(hashedBytes[:])
	return hashedString
}
