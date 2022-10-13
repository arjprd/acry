package algo

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/arjprd/crypt-service/driver"
)

// type AlgoTest struct {
// 	algoName    string
// 	algoHandler func()
// }

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!@#$&*./")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestBcryptDrive(t *testing.T) {
	fmt.Println("bcrypt algo")
	config := driver.NewConfig("../test/config.yaml")
	for cost := 4; cost < 12; cost++ {
		password := RandStringRunes(rand.Intn(32))
		handler := NewBcryptHash(cost, config)
		hash, err := handler.Generate(password)
		if err != nil {
			t.Error("error on bcrypt hash generation")
			return
		}
		if !handler.Verify(hash, password) {
			t.Error("error on bcrypt hash verification")
			return
		}
	}
}

func TestPbkdf2Drive(t *testing.T) {
	fmt.Println("pbkdf2 algo")
	config := driver.NewConfig("../test/config.yaml")
	for _, hashFunc := range []string{"sha1", "md5", "sha256", "sha512"} {
		password := RandStringRunes(rand.Intn(32))
		salt := RandStringRunes(rand.Intn(32))
		iter := rand.Intn(4999) + 1
		keylen := rand.Intn(32)
		handler := NewPbkdf2Hash(iter, salt, keylen, hashFunc, config)
		hash, err := handler.Generate(password)
		if err != nil {
			t.Error("error on bcrypt hash generation")
			return
		}
		if !handler.Verify(hash, password) {
			t.Error("error on bcrypt hash verification")
			return
		}
	}
}
