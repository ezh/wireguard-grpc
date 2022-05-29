// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package logger_test

import (
	"math"
	"testing"

	"github.com/ezh/wireguard-grpc/internal/logger"
	"github.com/stretchr/testify/assert"
)

func TestMarshalVerbosity(t *testing.T) {
	testCases := []struct {
		v    logger.Verbosity
		name string
	}{
		{logger.DebugVerbosity, logger.VerbosityNameDebug},
		{logger.InfoVerbosity, logger.VerbosityNameInfo},
		{logger.SuccessVerbosity, logger.VerbosityNameSuccess},
		{logger.ErrorVerbosity, logger.VerbosityNameError},
		{logger.FatalVerbosity, logger.VerbosityNameFatal},
		{logger.SuppressVerbosity, logger.VerbosityNameSuppress},
		{logger.WarnVerbosity, logger.VerbosityNameWarn},
	}

	for _, tc := range testCases {
		n, err := tc.v.MarshalText()
		if err != nil {
			assert.Fail(t, "marshaling verbosity %q name: %v", tc.name, err)
		}
		assert.Equalf(t, string(n), tc.name, "expected %q, but got %q", tc.name, n)
	}
}

func TestMarshalVerbosityInvalid(t *testing.T) {
	vInvalid := logger.Verbosity(math.MaxInt32)
	_, err := vInvalid.MarshalText()
	assert.NotNilf(t, err, "must fail to marshal invalid verbosity level: %v", err)
}

func TestParseVerbosity(t *testing.T) {
	testCases := []struct {
		v    logger.Verbosity
		name string
	}{
		{logger.DebugVerbosity, logger.VerbosityNameDebug},
		{logger.InfoVerbosity, logger.VerbosityNameInfo},
		{logger.SuccessVerbosity, logger.VerbosityNameSuccess},
		{logger.ErrorVerbosity, logger.VerbosityNameError},
		{logger.FatalVerbosity, logger.VerbosityNameFatal},
		{logger.SuppressVerbosity, logger.VerbosityNameSuppress},
		{logger.WarnVerbosity, logger.VerbosityNameWarn},
	}

	for _, tc := range testCases {
		v, err := logger.ParseVerbosity(tc.name)
		if err != nil {
			assert.Fail(t, "parsing %q verbosity name: %v", tc.name, err)
		}
		assert.Equalf(t, v, tc.v, "expected %q, but got %q", tc.v, v)
	}
}

func TestParseVerbosityInvalid(t *testing.T) {
	vName := "invalid"
	_, err := logger.ParseVerbosity(vName)
	assert.NotNilf(t, err, "must fail to parse invalid verbosity level name %q: %v", vName, err)
}

func TestUnmarshalVerbosity(t *testing.T) {
	vOld := logger.DebugVerbosity
	vNew := logger.ErrorVerbosity
	err := vOld.UnmarshalText([]byte(vNew.String()))
	assert.Nilf(t, err, "unmarshalling verbosity name %q", logger.ErrorVerbosity.String())
	assert.Equalf(t, vOld, vNew, "expected %q, but got %q", vNew, vOld)
}

func TestUnmarshalVerbosityInvalid(t *testing.T) {
	v := logger.DebugVerbosity
	invalidName := "invalid"
	err := v.UnmarshalText([]byte(invalidName))
	assert.NotNilf(t, err, "unmarshal error: %v", err)
}

func TestVerbosityUnknownName(t *testing.T) {
	assert.Equal(t, logger.Verbosity(math.MaxUint32).String(), logger.VerbosityNameUnknown)
}
