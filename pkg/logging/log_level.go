package logging

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

type LogLevel string

const (
	_DebugLevel LogLevel = "debug"
	_InfoLevel  LogLevel = "info"
)

// LogLevelの実体
var LoggingLevel LogLevel

func init() {
	// default
	LoggingLevel = _InfoLevel

	lEnv := os.Getenv("LOG_LEVEL")
	if lEnv == "" {
		LoggingLevel = _InfoLevel
	} else {
		// env設定されていればその値を使用
		LoggingLevel = LogLevel(lEnv)
	}
}

// ConvertLogrusLevel Convert to logrus log level.
func (l LogLevel) ConvertLogrusLevel() (logrus.Level, error) {
	switch l {
	case _DebugLevel:
		return logrus.DebugLevel, nil
	case _InfoLevel:
		return logrus.InfoLevel, nil
	default:
	}
	return logrus.InfoLevel, errors.New("undefined log level")
}
