package logger

import (
	"strings"

	"github.com/go-logr/logr"
)

// verbosity level
type LogLevel int

type LogBuilder func(logLevel LogLevel, options ...Option) (error, logr.Logger)

const (
	// logging is disabled
	LOG_OFF LogLevel = iota - 1
	// logging channel for general information
	LOG_INFO
	// logging channel for detailed information
	LOG_DEBUG
	// logging channel for verbose information about the current state
	LOG_TRACE
)

// ParseLogLevel convert the string representation into an instance of log level.
// the conversion is case insensitive. the default result is LOG_OFF
func ParseLogLevel(value string) LogLevel {
	switch strings.ToLower(value) {
	case "trace":
		fallthrough
	case "spam":
		return LOG_TRACE

	case "insight":
		fallthrough
	case "debug":
		return LOG_DEBUG

	case "info":
		fallthrough
	case "information":
		fallthrough
	case "informative":
		return LOG_INFO

	default:
		return 0
	}
}
