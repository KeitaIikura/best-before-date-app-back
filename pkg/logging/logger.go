package logging

var (
	runtimeCall = 2
	Default     *CustomLogrus
)

func init() {

	logLevel, _ := LoggingLevel.ConvertLogrusLevel()

	// TODO: 開発環境でdebugが表示できるように修正

	config := LogrusConfig{
		enableJSONFormatter: false,
		logLevel:            logLevel,
	}

	Default = NewLogrusLogger(config, runtimeCall)
}

func Debug(xRequestID string, args ...interface{}) {
	Default.debug(xRequestID, args)
}

func Info(xRequestID string, args ...interface{}) {
	Default.info(xRequestID, args)
}

func Error(xRequestID string, args ...interface{}) {
	Default.error(xRequestID, args)
}

func Fatal(xRequestID string, args ...interface{}) {
	Default.fatal(xRequestID, args)
}
