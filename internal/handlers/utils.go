package handlers

import (
	"net/http"
)

func (h *URLHandler) checkMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return false
	}

	return true
}

func (h *URLHandler) StatusBadRequest(w http.ResponseWriter, r *http.Request) bool {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad request"))
	return false
}

func (h *URLHandler) StatusServerError(w http.ResponseWriter, r *http.Request) bool {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Something went wrong"))
	return false
}
