package validator

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

var EmailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) SetError(key string, msg string) {
	v.Errors[key] = msg
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) Check(expr bool, key string, msg string) {
	if !expr {
		v.SetError(key, msg)
	}
}

func (v *Validator) NotEmpty(value string, key string, msg string) {
	if value == "" {
		v.SetError(key, msg)
	}
}

func (v *Validator) MinLength(value string, minLen int, key string) {
	if utf8.RuneCountInString(value) < minLen {
		v.SetError(key, fmt.Sprintf("%s minimum length: %d", key, minLen))
	}
}

func (v *Validator) MaxLength(value string, maxLen int, key string) {
	if utf8.RuneCountInString(value) > maxLen {
		v.SetError(key, fmt.Sprintf("%s maximum length: %d", key, maxLen))
	}
}

func (v *Validator) BetweenLength(value string, minLen int, maxLen int, key string) {
	if utf8.RuneCountInString(value) < minLen || utf8.RuneCountInString(value) > maxLen {
		v.SetError(key, fmt.Sprintf("%s length must be between %d-%d", key, minLen, maxLen))
	}
}

func (v *Validator) Min(value int, min int, key string) {
	if value < min {
		v.SetError(key, fmt.Sprintf("%s minimum: %d", key, min))
	}
}

func (v *Validator) Max(value int, max int, key string) {
	if value > max {
		v.SetError(key, fmt.Sprintf("%s maximum: %d", key, max))
	}
}

func (v *Validator) Between(value int, min int, max int, key string) {
	if value < min || value > max {
		v.SetError(key, fmt.Sprintf("%s must be between %d-%d", key, min, max))
	}
}

func (v *Validator) MatchesRegexp(value string, rx *regexp.Regexp, key string) {
	if !rx.MatchString(value) {
		v.SetError(key, fmt.Sprintf("%s is invalid", key))
	}
}

func In[T comparable](value T, values ...T) bool {
	for i := range values {
		if value == values[i] {
			return true
		}
	}
	return false
}
