package tcpservice

import (
	"encoding/json"
	"errors"
	"net"

	"github.com/arjprd/crypt-service/driver"
)

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

type Operations struct {
	c         *driver.Config
	opHandler map[string]OperationHandler
}

type OperationHandler func(Request, Response)

func (h *TCPService) handleRequest(c net.Conn, requestData []byte) {
	var request Request
	err := json.Unmarshal(requestData, request)
	if err != nil {
		h.c.Logger().Error("json convertion failed %+v", err)
		return
	}
	response := Response{
		c:       c,
		isError: false,
	}

	h.o.handle(request, response)
}

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

func (s *Response) Send() error {
	data, err := json.Marshal(s)
	if err != nil {
		return errors.New("json marshall failed")
	}
	data = append(data, '\r')
	s.c.Write(data)
	return nil
}

func (s *Response) SetError(err string) {
	s.errMessage = err
	s.isError = true
}

func (s *Response) SetHash(hash string) {
	s.hash = hash
	s.isError = false
}

func (s *Response) SetVerified(verified bool) {
	s.verfied = verified
	s.isError = false
}
