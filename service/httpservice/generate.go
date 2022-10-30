package httpservice

import (
	"net/http"
)

const (
	GENERATE_HASH_ROUTE     = "/hash/generate/"
	METHOD_POST             = "POST"
	METHOD_NOT_ALLOWED_TEXT = "405 - Method Not Allowed"
	INTERNAL_SERVER_TEXT    = "500 - Something Went Wrong"
)

func (h *HttpService) GenerateHash(w http.ResponseWriter, r *http.Request) {
	if r.Method == METHOD_POST {

		if err := r.ParseForm(); err != nil {
			h.c.Logger().Error("ParseForm() err: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(INTERNAL_SERVER_TEXT))
			return
		}

		password := r.FormValue("password")

		algorithmHandler, err := h.findAlgo(r, GENERATE_HASH_ROUTE)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(INTERNAL_SERVER_TEXT))
			return
		}

		hash, err := algorithmHandler.Generate(password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(INTERNAL_SERVER_TEXT))
			return
		}
		w.Write([]byte(hash))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(METHOD_NOT_ALLOWED_TEXT))
	}
}
