package profiler

import (
	"testing"

	"mh-api/pkg/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStartPyroscope_SkipsWhenAddressEmpty(t *testing.T) {
	cfg := &config.Config{
		PyroscopeServerAddress: "",
		PyroscopeAPIKey:        "",
		Env:                    "test",
	}

	stop := StartPyroscope(cfg, "test-app")
	require.NotNil(t, stop)
	assert.NotPanics(t, func() { stop() })
}

func TestStartPyroscope_ReturnsStopFunctionOnInvalidAddress(t *testing.T) {
	cfg := &config.Config{
		PyroscopeServerAddress: "http://invalid-pyroscope-host-that-does-not-exist:4040",
		PyroscopeAPIKey:        "",
		Env:                    "test",
	}

	stop := StartPyroscope(cfg, "test-app")
	require.NotNil(t, stop)
	assert.NotPanics(t, func() { stop() })
}
