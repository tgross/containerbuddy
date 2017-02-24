package tests

import (
	"strings"
	"testing"
)

func TestEnvVars(t *testing.T) {
	out := composeRun(t, "testdata/TestEnvVars/compose.yml", "shell")
	if strings.Contains(out, "FAIL") {
		t.Fatalf("expected IP address, got FAIL")
	}
}
