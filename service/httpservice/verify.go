package httpservice

import (
	"net/http"
)

const (
	VERIFY_HASH_ROUTE = "/hash/verify/"
)

func (h *HttpService) VerifyHash(w http.ResponseWriter, r *http.Request) {
	if r.Method == METHOD_POST {

		if err := r.ParseForm(); err != nil {
			h.c.Logger().Error("ParseForm() err: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(INTERNAL_SERVER_TEXT))
			return
		}

		password := r.FormValue("password")
		hash := r.FormValue("hash")

		algorithmHandler, err := h.findAlgo(r, VERIFY_HASH_ROUTE)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(INTERNAL_SERVER_TEXT))
			return
		}

		match := algorithmHandler.Verify(hash, password)
		if match {
			w.Write([]byte("true"))
			return
		}
		w.Write([]byte("false"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(METHOD_NOT_ALLOWED_TEXT))
	}
}
