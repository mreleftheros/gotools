package context

import (
	"context"
	"fmt"
	"net/http"
)

type ContextKey string

func Set(r *http.Request, key ContextKey, value any) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), key, value))
}

func Get[T any](r *http.Request, key ContextKey) *T {
	v, ok := r.Context().Value(key).(*T)
	if !ok {
		panic(fmt.Sprintf("missing %s context value", key))
	}
	return v
}
