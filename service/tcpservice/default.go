package tcpservice

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/arjprd/acry/algo"
)

func (o *Operations) DefaultHandler(req Request, res Response) {
	err := res.Send()
	if err != nil {
		o.c.Logger().Error("unknown operation")
	}
}

func (h *TCPService) findAlgo(r Request) (algo.HashAlgorithm, error) {
	switch strings.ToLower(r.algorithm) {
	case algo.ALGO_NAME_BCRYPT:
		cost := algo.DEFAULT_BCRYPT_COST
		var param struct {
			Cost int `json:"cost"`
		}
		var byteData []byte
		r.parameters.UnmarshalJSON(byteData)
		json.Unmarshal(byteData, param)
		cost = param.Cost
		return algo.NewBcryptHash(cost, h.c), nil
	case algo.ALGO_NAME_PBKDF2:
		iter := algo.DEFAULT_PBKDF2_ITER
		keylen := algo.DEFAULT_PBKDF2_KEYLEN
		hashFunc := algo.DEFAULT_PBKDF2_HASH_FUNC

		var param struct {
			Iter     int    `json:"iter"`
			Keylen   int    `json:"klen"`
			HashFunc string `json:"hf"`
			Salt     string `json:"salt"`
		}
		var byteData []byte
		r.parameters.UnmarshalJSON(byteData)
		json.Unmarshal(byteData, param)

		salt := param.Salt
		iter = param.Iter
		keylen = param.Keylen
		hashFunc = param.HashFunc

		return algo.NewPbkdf2Hash(iter, salt, keylen, hashFunc, h.c), nil
	case algo.ALGO_NAME_ARGON2I:
		time := algo.DEFAULT_ARGON2I_TIME
		keylen := algo.DEFAULT_ARGON2I_KEYLEN
		memory := algo.DEFAULT_ARGON2I_MEMORY
		threads := algo.DEFAULT_ARGON2I_THREADS

		var param struct {
			Time    int    `json:"time"`
			Keylen  int    `json:"klen"`
			Memory  int    `json:"mem"`
			Threads int    `json:"th"`
			Salt    string `json:"salt"`
		}
		var byteData []byte
		r.parameters.UnmarshalJSON(byteData)
		json.Unmarshal(byteData, param)

		salt := param.Salt
		time = param.Time
		keylen = param.Keylen
		memory = param.Memory
		threads = param.Threads

		return algo.NewArgon2iHash(salt, uint(time), uint32(memory), uint8(threads), uint32(keylen), h.c), nil
	case algo.ALGO_NAME_ARGON2ID:
		time := algo.DEFAULT_ARGON2I_TIME
		keylen := algo.DEFAULT_ARGON2I_KEYLEN
		memory := algo.DEFAULT_ARGON2I_MEMORY
		threads := algo.DEFAULT_ARGON2I_THREADS

		var param struct {
			Time    int    `json:"time"`
			Keylen  int    `json:"klen"`
			Memory  int    `json:"mem"`
			Threads int    `json:"th"`
			Salt    string `json:"salt"`
		}
		var byteData []byte
		r.parameters.UnmarshalJSON(byteData)
		json.Unmarshal(byteData, param)

		salt := param.Salt
		time = param.Time
		keylen = param.Keylen
		memory = param.Memory
		threads = param.Threads
		return algo.NewArgon2idHash(salt, uint(time), uint32(memory), uint8(threads), uint32(keylen), h.c), nil
	}

	h.c.Logger().Error("invalid algorithm: %s", r.algorithm)
	return nil, errors.New("invalid algorithm")
}
