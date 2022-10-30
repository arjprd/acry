package tcpservice

import (
	"encoding/json"
	"errors"
	"net"

	"github.com/arjprd/crypt-service/driver"
)

type Request struct {
	Algorithm  string          `json:"algo"`
	Operation  string          `json:"op"`
	Password   string          `json:"pass"`
	Hash       string          `json:"hash"`
	Parameters json.RawMessage `json:"param"`
}

type Response struct {
	c          net.Conn `json:"-"`
	Hash       string   `json:"hash"`
	Verfied    bool     `json:"verified"`
	IsError    bool     `json:"error"`
	ErrMessage string   `json:"message"`
}

type Operations struct {
	c         *driver.Config
	opHandler map[string]OperationHandler
}

type OperationHandler func(*Request, *Response)

func (h *TCPService) handleRequest(c net.Conn, requestData []byte) {
	request := &Request{}
	err := json.Unmarshal(requestData, request)
	if err != nil {
		h.c.Logger().Error("handleRequest:json convertion failed %+v", err)
		return
	}
	response := &Response{
		c:       c,
		IsError: false,
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

func (o *Operations) handle(req *Request, res *Response) {
	if handler, ok := o.opHandler[req.Operation]; ok {
		handler(req, res)
		return
	}
	o.DefaultHandler(req, res)
}

func (s *Response) Send() error {
	data, err := json.Marshal(s)
	if err != nil {
		return errors.New("send:json marshall failed")
	}
	data = append(data, '\r')
	s.c.Write(data)
	return nil
}

func (s *Response) SetError(err string) {
	s.ErrMessage = err
	s.IsError = true
}

func (s *Response) SetHash(hash string) {
	s.Hash = hash
	s.IsError = false
}

func (s *Response) SetVerified(verified bool) {
	s.Verfied = verified
	s.IsError = false
}
