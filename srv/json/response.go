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

func Write(l *slog.Logger, w http.ResponseWriter, r *http.Request, status int, data any, headers http.Header) {
	json, err := json.Marshal(data)
	if err != nil {
		WriteInternalError(l, w, r, err, nil)
		return
	}

	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(json)
	if err != nil {
		WriteInternalError(l, w, r, err, nil)
	}
}

func WriteInternalError(l *slog.Logger, w http.ResponseWriter, r *http.Request, err error, headers http.Header) {
	logError(l, r, err)
	data := NewErrorResponse(http.StatusText(http.StatusInternalServerError), nil)
	// Write(w, 500, NewErrorResponse(http.StatusText(http.StatusInternalServerError), nil), headers)
	json, err := json.Marshal(data)
	if err != nil {
		logError(l, r, err)
	}

	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_, err = w.Write(json)
	if err != nil {
		logError(l, r, err)
	}
}
