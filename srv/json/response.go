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

func logError(r *http.Request, err error) {
	slog.Error(err.Error(), slog.String("method", r.Method), slog.String("path", r.URL.Path))
}

func Write(w http.ResponseWriter, r *http.Request, status int, data any) {
	json, err := json.Marshal(data)
	if err != nil {
		WriteInternalError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(json)
	if err != nil {
		WriteInternalError(w, r, err)
	}
}

func WriteInternalError(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, err)
	data := NewErrorResponse(http.StatusText(http.StatusInternalServerError), nil)
	// Write(w, 500, NewErrorResponse(http.StatusText(http.StatusInternalServerError), nil), headers)
	json, err := json.Marshal(data)
	if err != nil {
		logError(r, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_, err = w.Write(json)
	if err != nil {
		logError(r, err)
	}
}
