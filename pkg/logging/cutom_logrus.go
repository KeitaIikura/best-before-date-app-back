package logging

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type CustomLogrus struct {
	logger      *logrus.Logger
	runtimeCall int
}

type LogrusConfig struct {
	enableJSONFormatter bool
	logLevel            logrus.Level
}

func NewLogrusLogger(conf LogrusConfig, runtimeCall int) *CustomLogrus {

	lgrs := logrus.New()

	lgrs.SetReportCaller(false)
	if conf.enableJSONFormatter {
		lgrs.SetFormatter(&logrus.JSONFormatter{DisableHTMLEscape: true})
	} else {
		lgrs.SetFormatter(&logrus.TextFormatter{})
	}
	lgrs.Level = conf.logLevel

	cLogrus := &CustomLogrus{
		logger:      lgrs,
		runtimeCall: runtimeCall,
	}

	return cLogrus
}

func (cl *CustomLogrus) debug(xRequestID string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(cl.runtimeCall)
	cl.logger.WithFields(logrus.Fields{
		"x-request-id": xRequestID,
		"file":         fmt.Sprintf("%s:%d", file, line),
	}).Log(logrus.DebugLevel, strings.ReplaceAll(fmt.Sprintf("%+v", args...), "\"", ""))
}

func (cl *CustomLogrus) info(xRequestID string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(cl.runtimeCall)
	cl.logger.WithFields(logrus.Fields{
		"x-request-id": xRequestID,
		"file":         fmt.Sprintf("%s:%d", file, line),
	}).Log(logrus.InfoLevel, strings.ReplaceAll(fmt.Sprintf("%+v", args...), "\"", ""))
}

func (cl *CustomLogrus) error(xRequestID string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(cl.runtimeCall)
	cl.logger.WithFields(logrus.Fields{
		"x-request-id": xRequestID,
		"file":         fmt.Sprintf("%s:%d", file, line),
	}).Log(logrus.ErrorLevel, strings.ReplaceAll(fmt.Sprintf("%+v", args...), "\"", ""))
}

func (cl *CustomLogrus) fatal(xRequestID string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(cl.runtimeCall)
	cl.logger.WithFields(logrus.Fields{
		"x-request-id": xRequestID,
		"file":         fmt.Sprintf("%s:%d", file, line),
	}).Fatal(strings.ReplaceAll(fmt.Sprintf("%+v", args...), "\"", ""))
}
