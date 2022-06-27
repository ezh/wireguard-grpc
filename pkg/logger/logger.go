package logger

import (
	"strings"

	"github.com/go-logr/logr"
)

// verbosity level
type LogLevel int

type LogBuilder func(logLevel LogLevel, options ...Option) (logr.Logger, error)

const (
	// logging is disabled
	LogOff LogLevel = iota - 1
	// logging channel for general information
	LogInfo
	// logging channel for detailed information
	LogDebug
	// logging channel for verbose information about the current state
	LogTrace
)

// ParseLogLevel convert the string representation into an instance of log level.
// the conversion is case insensitive. the default result is LOG_OFF
func ParseLogLevel(value string) LogLevel {
	switch strings.ToLower(value) {
	case "trace":
		fallthrough
	case "spam":
		return LogTrace

	case "insight":
		fallthrough
	case "debug":
		return LogDebug

	case "info":
		fallthrough
	case "information":
		fallthrough
	case "informative":
		return LogInfo

	default:
		return 0
	}
}

func NewLogger(logBuilder LogBuilder, rawLogLevel string, verbosity int) (logr.Logger, error) {
	if verbosity == -1 {
		logBuilder = func(logLevel LogLevel, options ...Option) (logr.Logger, error) {
			return logr.Discard(), nil
		}
	}
	logLevel := ParseLogLevel(rawLogLevel)
	return logBuilder(LogLevel(int(logLevel) + verbosity))
}
