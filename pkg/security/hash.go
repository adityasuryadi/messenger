package security

import (
	"crypto/sha1"
	"encoding/hex"
)

func Hash(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
