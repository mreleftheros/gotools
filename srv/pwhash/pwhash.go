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

type hashParams struct {
	argon2Version int
	hashBytes     []byte
	saltBytes     []byte
	t             uint32 // times
	m             uint32 // memory
	p             uint8  // parallelism
	keyLen        uint32
}

func newHashParams() *hashParams {
	return &hashParams{
		argon2Version: argon2.Version,
		t:             2,
		m:             19 * 1024,
		p:             1,
		keyLen:        32,
	}
}

func Hash(password string) (string, error) {
	hp := newHashParams()
	saltBytes := make([]byte, 16)
	if _, err := rand.Read(saltBytes); err != nil {
		return "", err
	}
	hp.saltBytes = saltBytes
	hp.hashBytes = argon2.IDKey([]byte(password), hp.saltBytes, hp.t, hp.m, hp.p, hp.keyLen)

	return fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		hp.argon2Version,
		hp.m,
		hp.t,
		hp.p,
		base64.RawStdEncoding.EncodeToString(hp.saltBytes),
		base64.RawStdEncoding.EncodeToString(hp.hashBytes),
	), nil
}

func Verify(h, password string) (bool, error) {
	hp, err := parseHash(h)
	if err != nil {
		return false, err
	}

	hash := argon2.IDKey(
		[]byte(password),
		hp.saltBytes,
		hp.t,
		hp.m,
		hp.p,
		hp.keyLen,
	)

	return subtle.ConstantTimeCompare(hp.hashBytes, hash) == 1, nil
}

func parseHash(h string) (params *hashParams, err error) {
	vals := strings.Split(h, "$")
	if len(vals) != 6 {
		return nil, errors.New("invalid hash format")
	}

	hp := &hashParams{}
	if !strings.HasPrefix(vals[1], "argon2id") {
		return nil, errors.New("unsupported algorithm")
	}
	fmt.Sscanf(vals[2], "v=%d", &hp.argon2Version)

	fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &hp.m, &hp.t, &hp.p)

	saltBytes, err := base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, err
	}
	hp.saltBytes = saltBytes

	hashBytes, err := base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, err
	}
	hp.hashBytes = hashBytes
	hp.keyLen = uint32(len(hashBytes))

	return hp, nil
}
