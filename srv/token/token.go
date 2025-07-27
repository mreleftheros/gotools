package token

import (
	"crypto/rand"
	"encoding/base64"
)

const DEFAULT_LENGTH = 16

type TokenOpts struct {
	Length int
}

func New(opts *TokenOpts) (string, error) {
	if opts == nil {
		opts = &TokenOpts{
			Length: DEFAULT_LENGTH,
		}
	}

	buf := make([]byte, opts.Length)

	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}

	tok := base64.RawStdEncoding.WithPadding(base64.NoPadding).EncodeToString(buf)

	return tok, nil
}
