package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	err := New("./sample/")
	require.NoError(t, err)
}
