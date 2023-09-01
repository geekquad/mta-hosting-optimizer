package service

import (
	"os"
	"testing"
)

func TestGetThresholdFromEnv_CustomValue(t *testing.T) {
	os.Setenv("THRESHOLD", "5")

	threshold := GetThresholdFromEnv()

	if threshold != 5 {
		t.Errorf("Expected threshold to be 5, but got %d", threshold)
	}

	os.Unsetenv("THRESHOLD")
}

func TestGetThresholdFromEnv_DefaultValue(t *testing.T) {
	os.Unsetenv("THRESHOLD")
	threshold := GetThresholdFromEnv()

	if threshold != 1 {
		t.Errorf("Expected threshold to be 1, but got %d", threshold)
	}
}
