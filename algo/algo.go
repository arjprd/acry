package algo

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

type HashAlgorithm interface {
	// generate hash in string format, converted to base64 if not bytes not in readable string range
	Generate(password string) (hash string, err error)
	// verifies the provided password and hash
	Verify(hash string, password string) (match bool)
}

func getHashAlgorithm(h string) func() hash.Hash {
	switch h {
	case "sha1":
		return sha1.New
	case "md5":
		return md5.New
	case "sha256":
		return sha256.New
	case "sha512":
		return sha512.New
	}
	return sha1.New
}
