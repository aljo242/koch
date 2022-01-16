package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	err := New("./sample/")
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
