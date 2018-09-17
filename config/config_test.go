package config

import (
	"os"
	"testing"
)

func TestNew_NoEnvironmentVariables(t *testing.T) {
	os.Clearenv()
	_, err := New()

	if err == nil {
		t.Error("should fail if env variables are not existing")
	}
}

func TestNew_Env(t *testing.T) {
	os.Clearenv()
	os.Setenv("PORT", "1234")
	os.Setenv("SERVER_HOST", "test")
	os.Setenv("SERVICE_TYPE", "client")

	c, err := New()
	if err != nil {
		t.Errorf("failed to load config: %v", err)
	}

	if c.Port != "1234" {
		t.Error("Failed to load PORT environment variable")
	}
}
