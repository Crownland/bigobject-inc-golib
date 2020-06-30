package logger

// Logger logger
type Logger interface {
	Fatal(v ...interface{})
	Fatalf(format string, args ...interface{})
	Error(v ...interface{})
	Errorf(format string, args ...interface{})
	Warn(v ...interface{})
	Warnf(format string, args ...interface{})
	Info(v ...interface{})
	Infof(format string, args ...interface{})
	Debug(v ...interface{})
	Debugf(format string, args ...interface{})
}

// LevelFatal logger level: fatal
const LevelFatal = "fatal"

// LevelError logger level: error
const LevelError = "error"

// LevelWarn logger level: warn
const LevelWarn = "warn"

// LevelInfo logger level: info
const LevelInfo = "info"

// LevelDebug logger level: debug
const LevelDebug = "debug"

// LevelDisable logger level: disable
const LevelDisable = "disable"

// Defer logger defer call
type Defer func()
