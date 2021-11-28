package ip_util

import (
	"testing"
)

// TestExternalIP verifies that TestExternalIP functions without returning an error
func TestExternalIP(t *testing.T) {
	_, err := ExternalIP()
	if err != nil {
		t.Errorf("ExternalIP failed : %w", err)
	}
}

// TestHostInfo() verifies that Host-functions without returning error
func TestHost(t *testing.T) {
	_, err := HostInfo()
	if err != nil {
		t.Errorf("HostInfo failed : %w", err)

	}
}
