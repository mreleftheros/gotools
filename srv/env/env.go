package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseString(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func ParseBool(key string, fallback bool) bool {
	v := os.Getenv(key)
	switch v {
	case "":
		return fallback
	case "true":
		return true
	case "false":
		return false
	default:
		panic(fmt.Errorf("environment key %s is invalid bool", key))
	}
}

func ParseInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return i
}

func ParseCSV(key string, fallback []string) []string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return strings.Split(v, ",")
}
