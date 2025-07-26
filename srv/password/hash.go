package password

import (
	"crypto/subtle"

	"golang.org/x/crypto/argon2"
)

func Hash(password string, salt string) []byte {
	return argon2.IDKey([]byte(password), []byte(salt), 2, 19*1024, 1, 32)
}

func Compare(provided string, salt string, stored []byte) bool {
	h := Hash(provided, salt)
	return subtle.ConstantTimeCompare(h, stored) == 1
}
