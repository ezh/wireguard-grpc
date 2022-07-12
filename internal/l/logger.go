package l

import (
	"log"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
)

var l *logr.Logger

func init() {
	logger := stdr.New(log.Default())
	l = &logger
}

// Info logs a non-error message with the given key/value pairs as context.
//
// The msg argument should be used to add some constant description to the log
// line.  The key/value pairs can then be used to add additional variable
// information.  The key/value pairs must alternate string keys and arbitrary
// values.
func Info(msg string, keysAndValues ...interface{}) {
	l.Info(msg, keysAndValues...)
}

// Error logs an error, with the given message and key/value pairs as context.
// It functions similarly to Info, but may have unique behavior, and should be
// preferred for logging errors (see the package documentations for more
// information). The log message will always be emitted, regardless of
// verbosity level.
//
// The msg argument should be used to add context to any underlying error,
// while the err argument should be used to attach the actual error that
// triggered this log line, if present. The err parameter is optional
// and nil may be passed instead of an error instance.
func Error(err error, msg string, keysAndValues ...interface{}) {
	l.Error(err, msg, keysAndValues...)
}

// V returns a new Logger instance for a specific verbosity level, relative to
// this Logger.  In other words, V-levels are additive.  A higher verbosity
// level means a log message is less important.  Negative V-levels are treated
// as 0.
func V(level int) logr.Logger {
	return l.V(level)
}

// RegisterLogger registers global logr implementation
func RegisterLogger(ll *logr.Logger) {
	l = ll
}
