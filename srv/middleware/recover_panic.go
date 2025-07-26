package middleware

import (
	"net/http"

	"github.com/mreleftheros/gotools/srv/json"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				json.WriteInternalError(w, r, err.(error))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
