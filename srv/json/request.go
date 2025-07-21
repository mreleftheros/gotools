package json

import (
	"encoding/json"
	"net/http"
)

// Parse parses json body and, if is successful, decodes the provided data to dest which must be a pointer.
func Parse(w http.ResponseWriter, r *http.Request, dest any) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1_048_576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dest)
	if err != nil {
		return err
	}

	return nil
}
