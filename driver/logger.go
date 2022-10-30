package driver

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Logger struct {
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
}

func NewLogger() *Logger {
	logger := &Logger{}
	if viper.GetBool(VIPER_KEY_SERVICE_INFO_LOG) {
		logger.infoLogger = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lmicroseconds)
	}
	if viper.GetBool(VIPER_KEY_SERVICE_ERROR_LOG) {
		logger.errorLogger = log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime|log.Lmicroseconds)
	}
	if viper.GetBool(VIPER_KEY_SERVICE_WARNING_LOG) {
		logger.warningLogger = log.New(os.Stdout, "WARNING ", log.Ldate|log.Ltime|log.Lmicroseconds)
	}
	return logger
}

func (l *Logger) Error(format string, v ...any) {
	if l.errorLogger != nil {
		l.errorLogger.Printf(format, v...)
	}
}

func (l *Logger) Warning(format string, v ...any) {
	if l.warningLogger != nil {
		l.warningLogger.Printf(format, v...)
	}
}

func (l *Logger) Info(format string, v ...any) {
	if l.infoLogger != nil {
		l.infoLogger.Printf(format, v...)
	}
}
