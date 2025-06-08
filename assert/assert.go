package assert

import (
	"strings"
	"testing"
)

type Case struct {
	Name     string
	Args     map[string]any
	Expected []any
}

func Equal[T comparable](t *testing.T, expected T, actual T) {
	t.Helper()

	if expected != actual {
		t.Errorf("expected: %v, actual: %v\n", expected, actual)
	}
}

func StringContains(t *testing.T, actual string, expectedSubstring string) {
	t.Helper()

	if !strings.Contains(actual, expectedSubstring) {
		t.Errorf("string \"%q\" expected to contain: %q\n", actual, expectedSubstring)
	}
}
