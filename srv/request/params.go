package request

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/mreleftheros/gotools/srv/validator"
)

// ParsePathID parses the id value from the request path
func ParsePathID(r *http.Request) (int64, error) {
	idParam, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || idParam < 1 {
		return 0, errors.New("invalid id path value")
	}
	return idParam, nil
}

// ParseQueryString parses a string from a query key or returns default value.
func ParseQueryString(qs url.Values, key string, defaultValue string) string {
	v := qs.Get(key)
	if v == "" {
		return defaultValue
	}
	return strings.ToLower(v)
}

// ParseQueryInt parses key from query string and returns int or a default value if fails.
func ParseQueryInt(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
	val := qs.Get(key)
	if val == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		v.SetError(key, fmt.Sprintf("%s query parameter must be an integer", key))
	}

	return i
}

// ParseQueryCSV parses comma seperated values in key and returns them on a slice or returns default value.
func ParseQueryCSV(qs url.Values, key string, defaultValue []string) []string {
	v := qs.Get(key)
	if v == "" {
		return defaultValue
	}
	return strings.Split(strings.ToLower(v), ",")
}
