package validator

import (
	"testing"

	"github.com/mreleftheros/gotools/assert"
)

func TestValidator(t *testing.T) {
	t.Parallel()

	_ = NewValidator()
}

func TestValidatorSetError(t *testing.T) {
	v := NewValidator()
	v.SetError("name", "Name cannot be empty")
	if vLen := len(v.Errors); vLen != 1 {
		t.Errorf("expected length: %d, actual: %d\n", 1, vLen)
	}
	assert.Equal(t, "Name cannot be empty", v.Errors["name"])
}

func TestValidatorEmpty(t *testing.T) {
	t.Parallel()
}
