package request

import (
	"net/url"
	"strconv"
	"strings"
)

func ParseQueryString(qs url.Values, key string, defaultValue string) string {
	v := qs.Get(key)
	if v == "" {
		return defaultValue
	}
	return strings.ToLower(v)
}

func ParseQueryInt(qs url.Values, key string, defaultValue int) int {
	val := qs.Get(key)
	if val == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}

	return i
}

func ParseQueryCSV(qs url.Values, key string, defaultValue []string) []string {
	v := qs.Get(key)
	if v == "" {
		return defaultValue
	}
	return strings.Split(strings.ToLower(v), ",")
}
