package helper

import (
	"crypto/sha1"
	"encoding/hex"
)

func Encrypt(Value any) string {
	hasher := sha1.New()
	hasher.Write([]byte(Value.(string)))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
