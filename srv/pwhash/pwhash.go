package pwhash

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type HashParams struct {
	Argon2Version int
	Hash          []byte
	Salt          []byte
	T             uint32 // times
	M             uint32 // memory
	P             uint8  // threads
	KeyLen        uint32
}

func Hash(password string) (string, error) {
	hp := &HashParams{
		Argon2Version: argon2.Version,
		T:             2,
		M:             19 * 1024,
		P:             1,
		KeyLen:        32,
	}

	salt, err := genSalt(16)
	if err != nil {
		return "", err
	}
	hp.Salt = salt
	hp.Hash = argon2.IDKey([]byte(password), hp.Salt, hp.T, hp.M, hp.P, hp.KeyLen)

	return fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		hp.Argon2Version,
		hp.M,
		hp.T,
		hp.P,
		base64.RawStdEncoding.EncodeToString(hp.Salt),
		base64.RawStdEncoding.EncodeToString(hp.Hash),
	), nil
}

func Verify(storedHash, password string) (bool, error) {
	hp, err := parseHash(storedHash)
	if err != nil {
		return false, err
	}

	hash := argon2.IDKey(
		[]byte(password),
		hp.Salt,
		hp.T,
		hp.M,
		hp.P,
		hp.KeyLen,
	)

	return subtle.ConstantTimeCompare(hp.Hash, hash) == 1, nil
}

func genSalt(size int) ([]byte, error) {
	buf := make([]byte, size)
	if _, err := rand.Read(buf); err != nil {
		return nil, err
	}

	return buf, nil
}

func parseHash(encodedHash string) (params *HashParams, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, errors.New("invalid hash format")
	}

	hp := &HashParams{}
	if !strings.HasPrefix(vals[1], "argon2id") {
		return nil, errors.New("unsupported algorithm")
	}
	fmt.Sscanf(vals[2], "v=%d", &hp.Argon2Version)

	fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &hp.M, &hp.T, &hp.P)

	salt, err := base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, err
	}
	hp.Salt = salt

	hash, err := base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, err
	}
	hp.Hash = hash
	hp.KeyLen = uint32(len(hash))

	return hp, nil
}
