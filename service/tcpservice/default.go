package tcpservice

import (
	"encoding/json"
	"net"

	"github.com/arjprd/crypt-service/driver"
)

type Operations struct {
	c         *driver.Config
	opHandler map[string]OperationHandler
}

type Request struct {
	algorithm  string          `json:"algo"`
	operation  string          `json:"op"`
	password   string          `json:"pass"`
	hash       string          `json:"hash"`
	parameters json.RawMessage `json:"param"`
}

type Response struct {
	c          net.Conn `json:"-"`
	hash       string   `json:"hash"`
	verfied    bool     `json:"verified"`
	isError    bool     `json:"error"`
	errMessage string   `json:"message"`
}

type OperationHandler func(Request, Response)

func NewOperations(c *driver.Config) *Operations {
	return &Operations{
		c:         c,
		opHandler: make(map[string]OperationHandler),
	}
}

func (o *Operations) RegisterOperation(operation string, handler OperationHandler) {
	o.opHandler[operation] = handler
}

func (o *Operations) handle(req Request, res Response) {
	if handler, ok := o.opHandler[req.algorithm]; ok {
		handler(req, res)
		return
	}
	o.DefaultHandler(req, res)
}

func (h *TCPService) handleRequest(c net.Conn, requestData []byte) {
	var request Request
	err := json.Unmarshal(requestData, request)
	if err != nil {
		h.c.Logger().Error("json convertion failed %+v", err)
		return
	}
	response := Response{
		c: c,
	}

	h.o.handle(request, response)
}

func (o *Operations) DefaultHandler(req Request, res Response) {

}
