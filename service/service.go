package service

import (
	"errors"

	"github.com/arjprd/acry/driver"
	"github.com/arjprd/acry/service/httpservice"
	"github.com/arjprd/acry/service/tcpservice"
)

func NewService(c *driver.Config) (driver.ServiceHandler, error) {
	return getService(c)
}

func getService(c *driver.Config) (driver.ServiceHandler, error) {
	switch c.GetServiceType() {
	case driver.SERVICE_TYPE_HTTP:
		return httpservice.NewHTTPService(c), nil
	case driver.SERVICE_TYPE_TCP:
		return tcpservice.NewTCPService(c), nil
	default:
		return nil, errors.New("invalid service type")
	}
}
