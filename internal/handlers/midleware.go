package handlers

import (
	"net/http"
)

type UserKey struct{}

func Middleware(h http.Handler) http.HandlerFunc {
	// TODO: проверка авторизации
	logFn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}

	return logFn
}
