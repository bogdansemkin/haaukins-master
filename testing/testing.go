package testing

import (
	"os"
	"testing"
)

func SkipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skipf("Ignore test %s in CI environment", t.Name())
	}
}

func SkipOnGh(t *testing.T) {
	if os.Getenv("GITHUB_CI") != "" {
		t.Skipf("Test %s ignored on GH ACTIONS CI", t.Name())
	}
}
