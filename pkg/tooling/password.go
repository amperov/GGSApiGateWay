package tooling

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

var salt = os.Getenv("HASH-SECRET")

func Hash(password string) string {

	hash := sha256.New()
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum([]byte(salt)))
}
