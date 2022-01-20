package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	err := New("./sample/")
	require.NoError(t, err)

	require.Equal(t, "debug", LogLevel())
	require.Equal(t, false, ServerSecure())
	require.Equal(t, "localhost", ServerIP())
	require.Equal(t, "80", ServerPort())
	require.Equal(t, false, ServerChooseIP())
	require.Equal(t, "", ServerCertFile())
	require.Equal(t, "", ServerKeyFile())
	require.Equal(t, "", ServerRootCA())
	require.Equal(t, "localhost", ServerHost())
	require.Equal(t, -3, ServerShutdownCode())
	require.Equal(t, false, ServerCmdEnable())
	require.Equal(t, 180, ServerCacheMaxAge())
	require.Equal(t, "myApp", AppName())
	require.Equal(t, "cozart shmoopler", OwnerName())
}

func TestNewNoPath(t *testing.T) {
	err := New("./incorrect/")
	require.NoError(t, err)
}

// TODO reimplement
/*
func TestConfig(t *testing.T) {
	// provide nonexistent file to get incorrect file error
	_, err := config.LoadConfig(incorrectConfigFile)
	require.ErrorIs(t, err, os.ErrNotExist)

	_, err = config.LoadConfig(sampleHTML)
	require.ErrorIs(t, err, config.ErrConfigNotJSON)
}
*/
