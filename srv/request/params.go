package request

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ParsePathID parses the id value from the request path
func ParsePathID(r *http.Request) (int64, error) {
	idParam, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || idParam < 1 {
		return 0, errors.New("invalid id path value")
	}
	return idParam, nil
}

type Query struct {
	url.Values
}

func NewQuery(values url.Values) *Query {
	return &Query{values}
}

func (q *Query) ParseString(key string, defaultValue string) string {
	v := q.Get(key)
	if v == "" {
		return defaultValue
	}
	return strings.ToLower(v)
}

func (q *Query) ParseInt(key string, defaultValue int) int {
	val := q.Get(key)
	if val == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}

	return i
}

func (q *Query) ParseCSV(key string, defaultValue []string) []string {
	v := q.Get(key)
	if v == "" {
		return defaultValue
	}
	return strings.Split(strings.ToLower(v), ",")
}
