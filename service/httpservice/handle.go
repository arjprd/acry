package httpservice

import (
	"net/http"
	"strconv"

	"github.com/arjprd/acry/driver"
)

type HttpService struct {
	c *driver.Config
}

func NewHTTPService(c *driver.Config) driver.ServiceHandler {
	return &HttpService{
		c: c,
	}
}

//scrypt, argon2
func (h *HttpService) RegisterAllRoutes() {
	http.HandleFunc(DEFAULT_ROUTE, h.All)
	http.HandleFunc(GENERATE_HASH_ROUTE, h.GenerateHash)
	http.HandleFunc(VERIFY_HASH_ROUTE, h.VerifyHash)
}

func (h *HttpService) Start() {
	port := h.c.GetServicePort()
	host := ":" + strconv.Itoa(port)
	h.RegisterAllRoutes()
	h.c.Logger().Info("HTTP service listening on port %d", port)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		h.c.Logger().Error("unable to start http server %+v", err)
	}
}
