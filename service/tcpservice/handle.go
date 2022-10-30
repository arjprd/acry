package tcpservice

import (
	"bufio"
	"net"
	"strconv"

	"github.com/arjprd/crypt-service/driver"
)

type TCPService struct {
	c *driver.Config
	o *Operations
}

func NewTCPService(c *driver.Config) driver.ServiceHandler {
	opHandler := NewOperations(c)
	serviceHandler := &TCPService{
		c: c,
		o: opHandler,
	}
	opHandler.RegisterOperation(TCP_OPERATION_GENERATION, serviceHandler.GenerateHandler)
	opHandler.RegisterOperation(TCP_OPERATION_VERIFY, serviceHandler.VerifyHandler)
	return serviceHandler
}

func (h *TCPService) Start() {
	port := h.c.GetServicePort()
	host := ":" + strconv.Itoa(port)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		h.c.Logger().Error("start:unable to start tcp server %+v", err)
	} else {
		h.c.Logger().Info("start:tcp service up on %+v", host)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			h.c.Logger().Error("start:something went wrong %+v", err)
		}
		go h.handler(conn)
	}
}

func (h *TCPService) handler(c net.Conn) {
	reader := bufio.NewReader(c)
	for {
		data, err := reader.ReadBytes('\r')
		if err != nil {
			h.c.Logger().Error("handler:error on reading data from connection %+v", err)
			break
		}
		h.handleRequest(c, data)
	}
}
