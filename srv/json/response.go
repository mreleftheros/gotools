package json

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type dataResponse struct {
	Data any `json:"data"`
}

func NewDataResponse(data any) dataResponse {
	return dataResponse{data}
}

type errorResponse struct {
	Error  string            `json:"error"`
	Errors map[string]string `json:"errors,omitempty"`
}

func NewErrorResponse(message string, errors map[string]string) errorResponse {
	return errorResponse{message, errors}
}

func logError(l *slog.Logger, r *http.Request, err error) {
	l.Error(err.Error(), "method", r.Method, "path", r.URL.Path)
}

func Write(w http.ResponseWriter, status int, data any, headers http.Header) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(json)
	if err != nil {
		return err
	}
	return nil
}

func WriteInternalError(l *slog.Logger, w http.ResponseWriter, r *http.Request, err error, headers http.Header) {
	logError(l, r, err)

	err = Write(w, http.StatusInternalServerError, NewErrorResponse(http.StatusText(http.StatusInternalServerError), nil), headers)
	if err != nil {
		logError(l, r, err)
	}
}
