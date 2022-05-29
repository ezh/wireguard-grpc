package logger

import "github.com/go-logr/logr"

type LogBuilder func(logLevel int, options ...Option) (error, logr.Logger)
