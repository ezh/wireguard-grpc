package logger

import (
	"io"
	"os"

	"go.uber.org/zap"
)

// logOptions contains predefined subset of superset of all possible logr loggers options
type logOptions struct {
	Output     io.Writer
	ZapOptions []zap.Option
}

// Option defines option function
type Option func(*logOptions)

// WithOutput sets output to io.Writer
func WithOutput(w io.Writer) Option {
	return func(o *logOptions) {
		o.Output = w
	}
}

// WithZapOptions adds Zap specific options to logger
func WithZapOptions(zo ...zap.Option) Option {
	return func(o *logOptions) {
		o.ZapOptions = append(o.ZapOptions, zo...)
	}
}

func Get(o ...Option) logOptions {
	lo := logOptions{
		Output: os.Stdout,
	}
	for _, fn := range o {
		fn(&lo)
	}
	return lo
}
