package request

import (
	"errors"
	"net/http"
	"strconv"
)

func ParsePathID(r *http.Request) (int64, error) {
	idParam, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || idParam < 1 {
		return 0, errors.New("invalid id param")
	}
	return idParam, nil
}
