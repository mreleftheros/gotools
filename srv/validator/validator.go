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

func NewValidator() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) SetError(key string, value string) {
	v.Errors[key] = value
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) Check(expr bool, key string, value string) {
	if !expr {
		v.SetError(key, value)
	}
}

func (v *Validator) NotEmpty(key string, s string) {
	if s == "" {
		v.SetError(key, fmt.Sprintf("%s cannot be empty", key))
	}
}

func (v *Validator) MinLength(key string, s string, n int) {
	if utf8.RuneCountInString(s) < n {
		v.SetError(key, fmt.Sprintf("%s must be minimum %d characters", key, n))
	}
}

func (v *Validator) MaxLength(key string, s string, n int) {
	if utf8.RuneCountInString(s) > n {
		v.SetError(key, fmt.Sprintf("%s must be maximum %d characters", key, n))
	}
}

func (v *Validator) BetweenLength(key string, s string, lower int, upper int) {
	if utf8.RuneCountInString(s) < lower || utf8.RuneCountInString(s) > upper {
		v.SetError(key, fmt.Sprintf("%s must be between %d-%d characters", key, lower, upper))
	}
}

// f
func (v *Validator) Min(key string, i int, l int) {
	if i < l {
		v.SetError(key, fmt.Sprintf("%s must be minimum %d", key, l))
	}
}

func (v *Validator) Max(key string, i int, l int) {
	if i > l {
		v.SetError(key, fmt.Sprintf("%s must be maximum %d", key, l))
	}
}

func (v *Validator) Between(key string, i int, lower int, upper int) {
	if i < lower || i > upper {
		v.SetError(key, fmt.Sprintf("%s must be between %d-%d", key, lower, upper))
	}
}

func (v *Validator) MatchesRegexp(key string, s string, rx *regexp.Regexp) {
	if !rx.MatchString(s) {
		v.SetError(key, fmt.Sprintf("%s is invalid", key))
	}
}

func In[T comparable](v T, values ...T) bool {
	for i := range values {
		if v == values[i] {
			return true
		}
	}
	return false
}
