package password

import (
	"crypto/subtle"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string, salt string) string {
	h := argon2.IDKey([]byte(password), []byte(salt), 2, 19*1024, 1, 32)
	return base64.RawStdEncoding.EncodeToString(h)
}

func ComparePassword(provided string, salt string, stored string) bool {
	h := HashPassword(provided, salt)
	return subtle.ConstantTimeCompare([]byte(h), []byte(stored)) == 1
}
