// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package l

import (
	"fmt"
	"strings"
)

// Verbosity level names.
const (
	VerbosityNameDebug    = "debug"
	VerbosityNameError    = "error"
	VerbosityNameFatal    = "fatal"
	VerbosityNameInfo     = "info"
	VerbosityNameSuccess  = "success"
	VerbosityNameSuppress = "suppress"
	VerbosityNameUnknown  = "unknown"
	VerbosityNameWarn     = "warn"
)

const (
	// FatalVerbosity is the verbosity level of the line printer for fatal messages.
	FatalVerbosity Verbosity = iota
	// ErrorVerbosity is the verbosity level of the line printer for error messages.
	ErrorVerbosity
	// WarnVerbosity is the verbosity level of the line printer for warning messages.
	WarnVerbosity
	// SuccessVerbosity is verbosity the level of the line printer for success messages.
	SuccessVerbosity
	// InfoVerbosity is verbosity the level of the line printer for information messages.
	InfoVerbosity
	// DebugVerbosity is verbosity the level of the line printer for debug messages.
	DebugVerbosity
	// SuppressVerbosity is the verbosity level of the line printer to suppress messages.
	SuppressVerbosity
)

// Verbosity defines the verbosity level of the line printer.
type Verbosity uint32

// ParseVerbosity takes a verbosity level name and returns the Verbosity level constant.
func ParseVerbosity(lvl string) (Verbosity, error) {
	switch strings.ToLower(lvl) {
	case VerbosityNameFatal:
		return FatalVerbosity, nil
	case VerbosityNameError:
		return ErrorVerbosity, nil
	case VerbosityNameWarn:
		return WarnVerbosity, nil
	case VerbosityNameSuccess:
		return SuccessVerbosity, nil
	case VerbosityNameInfo:
		return InfoVerbosity, nil
	case VerbosityNameDebug:
		return DebugVerbosity, nil
	case VerbosityNameSuppress:
		return SuppressVerbosity, nil
	}

	var v Verbosity
	return v, fmt.Errorf("not a valid level: %q", lvl)
}

// MarshalText returns the textual representation of itself.
func (v Verbosity) MarshalText() ([]byte, error) {
	switch v {
	case FatalVerbosity:
		return []byte(VerbosityNameFatal), nil
	case ErrorVerbosity:
		return []byte(VerbosityNameError), nil
	case WarnVerbosity:
		return []byte(VerbosityNameWarn), nil
	case SuccessVerbosity:
		return []byte(VerbosityNameSuccess), nil
	case InfoVerbosity:
		return []byte(VerbosityNameInfo), nil
	case DebugVerbosity:
		return []byte(VerbosityNameDebug), nil
	case SuppressVerbosity:
		return []byte(VerbosityNameSuppress), nil
	}

	return nil, fmt.Errorf("not a valid level %d", v)
}

// Convert the Verbosity to a string.
func (v Verbosity) String() string {
	if b, err := v.MarshalText(); err == nil {
		return string(b)
	}
	return VerbosityNameUnknown
}

// UnmarshalText implements encoding.TextUnmarshaler to unmarshal a textual representation of itself.
func (v *Verbosity) UnmarshalText(text []byte) error {
	l, err := ParseVerbosity(string(text))
	if err != nil {
		return err
	}

	*v = l
	return nil
}
