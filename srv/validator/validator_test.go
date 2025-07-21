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

func TestValidatorIsValid(t *testing.T) {
	testcases := []assert.Case{
		{
			Name: "valid",
			Args: []any{&Validator{
				Errors: map[string]string{},
			}},
			Expected: []any{0, true},
		},
		{
			Name: "invalid one",
			Args: []any{&Validator{
				Errors: map[string]string{
					"name": "Name cannot be empty",
				},
			}},
			Expected: []any{1, false},
		},
		{
			Name: "invalid two",
			Args: []any{&Validator{
				Errors: map[string]string{
					"name":     "Name cannot be empty",
					"password": "Password cannot be empty",
				},
			}},
			Expected: []any{2, false},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expected[0].(int), len(tc.Args[0].(*Validator).Errors))
			assert.Equal(t, tc.Expected[1].(bool), tc.Args[0].(*Validator).IsValid())
		})
	}
}

func TestValidatorEmpty(t *testing.T) {
	t.Parallel()

	testCases := [...]assert.Case{
		{
			Name:     "empty field valid",
			Args:     []any{""},
			Expected: []any{1, false},
		},
		{
			Name:     "empty field invalid",
			Args:     []any{"whatever"},
			Expected: []any{0, true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			v := NewValidator()
			v.Empty(tc.Args[0].(string), "name")

			assert.Equal(t, tc.Expected[0].(int), len(v.Errors))
			assert.Equal(t, tc.Expected[1].(bool), v.IsValid())
		})
	}
}
