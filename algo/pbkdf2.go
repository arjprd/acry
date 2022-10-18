package algo

import (
	"encoding/base64"
	"hash"

	"github.com/arjprd/acry/driver"
	"golang.org/x/crypto/pbkdf2"
)

const (
	ALGO_NAME_PBKDF2         = "pbkdf2"
	DEFAULT_PBKDF2_ITER      = 1
	DEFAULT_PBKDF2_KEYLEN    = 32
	DEFAULT_PBKDF2_HASH_FUNC = "sha1"
)

type Pbkdf2 struct {
	iter   int
	salt   []byte
	keyLen int
	h      func() hash.Hash
	c      *driver.Config
}

func NewPbkdf2Hash(iter int, salt string, keyLen int, h string, c *driver.Config) HashAlgorithm {
	return &Pbkdf2{
		iter:   iter,
		salt:   []byte(salt),
		keyLen: keyLen,
		h:      getHashAlgorithm(h),
		c:      c,
	}
}

func (a *Pbkdf2) Generate(password string) (string, error) {
	hash := pbkdf2.Key([]byte(password), a.salt, a.iter, a.keyLen, a.h)
	base64Hash := base64.StdEncoding.EncodeToString(hash)
	a.c.Logger().Info("pbkdf2 hash %+v", base64Hash)
	return base64Hash, nil
}

func (a *Pbkdf2) Verify(hash string, password string) bool {
	generatedhash, err := a.Generate(password)
	if err != nil {
		a.c.Logger().Error("pbkdf2 hash generation failed: %+v", err)
		return false
	}
	if generatedhash != hash {
		a.c.Logger().Error("pbkdf2 hash and password mismatch")
		return false
	}
	return true
}
