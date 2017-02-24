package tests

import (
	"strings"
	"testing"
)

func TestVersionFlag(t *testing.T) {
	out := composeRun(t, "testdata/TestVersionFlag/compose.yml", "shell")
	if !strings.Contains(out, "dev-build-not-for-release") {
		t.Fatalf("expected 'dev-build-not-for-release' in output, got %s", out)
	}
}
