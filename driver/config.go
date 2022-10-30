package driver

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type ServiceType string

type Config struct {
	port        int
	serviceType ServiceType
	logger      *Logger
}

var (
	SERVICE_TYPE_TCP  = ServiceType("TCP")
	SERVICE_TYPE_HTTP = ServiceType("HTTP")
)

var (
	VIPER_KEY_SERVICE_PORT        = "service.port"
	VIPER_KEY_SERVICE_TYPE        = "service.type"
	VIPER_KEY_SERVICE_INFO_LOG    = "log.info"
	VIPER_KEY_SERVICE_ERROR_LOG   = "log.error"
	VIPER_KEY_SERVICE_WARNING_LOG = "log.warning"
)

const (
	DEFAULT_SERVICE_PORT = 2909
)

func NewConfig(configPath string) *Config {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return &Config{
		port:        getServicePort(),
		serviceType: getServiceType(),
		logger:      NewLogger(),
	}
}

func getServiceType() ServiceType {
	return ServiceType(
		strings.ToUpper(viper.GetString(VIPER_KEY_SERVICE_TYPE)),
	)
}

func getServicePort() int {
	return viper.GetInt(VIPER_KEY_SERVICE_PORT)
}

func (c *Config) GetServiceType() ServiceType {
	return c.serviceType
}

func (c *Config) GetServicePort() int {
	if c.port == 0 {
		c.Logger().Warning("port not specified in config")
		return DEFAULT_SERVICE_PORT
	}
	return c.port
}

func (c *Config) Logger() *Logger {
	return c.logger
}
