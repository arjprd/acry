package algo

import (
	"github.com/arjprd/crypt-service/driver"
	"golang.org/x/crypto/bcrypt"
)

const (
	ALGO_NAME_BCRYPT    = "bcrypt"
	DEFAULT_BCRYPT_COST = bcrypt.DefaultCost
)

type Bcryt struct {
	cost int
	c    *driver.Config
}

func NewBcryptHash(cost int, c *driver.Config) HashAlgorithm {
	return &Bcryt{
		cost: cost,
		c:    c,
	}
}

func (a *Bcryt) Generate(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), a.cost)
	if err != nil {
		a.c.Logger().Error("bcrypt generate failsed %+v", err)
		return "", err
	}
	return string(hash), nil
}

func (a *Bcryt) Verify(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		a.c.Logger().Error("bcrypt verification failed: %+v", err)
		return false
	}
	return true
}
