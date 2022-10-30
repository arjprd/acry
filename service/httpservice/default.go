package httpservice

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/arjprd/crypt-service/algo"
)

const (
	DEFAULT_ROUTE = "/"
)

func (h *HttpService) All(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service Up")
}

func (h *HttpService) findAlgo(r *http.Request, prefix string) (algo.HashAlgorithm, error) {
	algorithm := strings.TrimPrefix(r.URL.Path, prefix)
	switch strings.ToLower(algorithm) {
	case algo.ALGO_NAME_BCRYPT:
		cost := algo.DEFAULT_BCRYPT_COST
		costInString := r.FormValue("cost")
		if costInString != "" {
			var err error
			cost, err = strconv.Atoi(costInString)
			if err != nil {
				h.c.Logger().Error("string to integer convertion err: %v", err)
				return nil, err
			}
		}
		return algo.NewBcryptHash(cost, h.c), nil
	case algo.ALGO_NAME_PBKDF2:
		iter := algo.DEFAULT_PBKDF2_ITER
		keylen := algo.DEFAULT_PBKDF2_KEYLEN
		hashFunc := algo.DEFAULT_PBKDF2_HASH_FUNC
		salt := r.FormValue("salt")

		if r.FormValue("iteration") != "" {
			var err error
			iter, err = strconv.Atoi(r.FormValue("iteration"))
			if err != nil {
				h.c.Logger().Error("string to integer convertion err: %v", err)
				return nil, err
			}
		}

		if r.FormValue("keylength") != "" {
			var err error
			keylen, err = strconv.Atoi(r.FormValue("keylength"))
			if err != nil {
				h.c.Logger().Error("string to integer convertion err: %v", err)
				return nil, err
			}
		}

		if r.FormValue("hashfunc") != "" {
			hashFunc = r.FormValue("hashfunc")
		}

		return algo.NewPbkdf2Hash(iter, salt, keylen, hashFunc, h.c), nil
	case algo.ALGO_NAME_ARGON2I:
		time := algo.DEFAULT_ARGON2I_TIME
		keylen := algo.DEFAULT_ARGON2I_KEYLEN
		memory := algo.DEFAULT_ARGON2I_MEMORY
		threads := algo.DEFAULT_ARGON2I_THREADS
		salt := r.FormValue("salt")

		if r.FormValue("time") != "" {
			var err error
			time, err = strconv.Atoi(r.FormValue("time"))
			if err != nil {
				h.c.Logger().Error("string to integer convertion err: %v", err)
				return nil, err
			}
		}

		if r.FormValue("keylength") != "" {
			var err error
			keylen, err = strconv.Atoi(r.FormValue("keylength"))
			if err != nil {
				h.c.Logger().Error("string to integer convertion err: %v", err)
				return nil, err
			}
		}

		if r.FormValue("memory") != "" {
			var err error
			memory, err = strconv.Atoi(r.FormValue("memory"))
			if err != nil {
				h.c.Logger().Error("string to integer convertion err: %v", err)
				return nil, err
			}
		}

		if r.FormValue("threads") != "" {
			var err error
			threads, err = strconv.Atoi(r.FormValue("threads"))
			if err != nil {
				h.c.Logger().Error("string to integer convertion err: %v", err)
				return nil, err
			}
		}

		return algo.NewArgon2iHash(salt, uint(time), uint32(memory), uint8(threads), uint32(keylen), h.c), nil
	case algo.ALGO_NAME_ARGON2ID:
		time := algo.DEFAULT_ARGON2ID_TIME
		keylen := algo.DEFAULT_ARGON2ID_KEYLEN
		memory := algo.DEFAULT_ARGON2ID_MEMORY
		threads := algo.DEFAULT_ARGON2ID_THREADS
		salt := r.FormValue("salt")

		if r.FormValue("time") != "" {
			var err error
			time, err = strconv.Atoi(r.FormValue("time"))
			if err != nil {
				h.c.Logger().Error("string to integer convertion err: %v", err)
				return nil, err
			}
		}

		if r.FormValue("keylength") != "" {
			var err error
			keylen, err = strconv.Atoi(r.FormValue("keylength"))
			if err != nil {
				h.c.Logger().Error("string to integer convertion err: %v", err)
				return nil, err
			}
		}

		if r.FormValue("memory") != "" {
			var err error
			memory, err = strconv.Atoi(r.FormValue("memory"))
			if err != nil {
				h.c.Logger().Error("string to integer convertion err: %v", err)
				return nil, err
			}
		}

		if r.FormValue("threads") != "" {
			var err error
			threads, err = strconv.Atoi(r.FormValue("threads"))
			if err != nil {
				h.c.Logger().Error("string to integer convertion err: %v", err)
				return nil, err
			}
		}

		return algo.NewArgon2idHash(salt, uint(time), uint32(memory), uint8(threads), uint32(keylen), h.c), nil
	}

	h.c.Logger().Error("invalid algorithm: %s", algorithm)
	return nil, errors.New("invalid algorithm")
}
