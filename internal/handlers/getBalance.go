package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kirsh-nat/gophermart.git/internal/app"
	userservices "github.com/kirsh-nat/gophermart.git/internal/services/userServices"
)

func (h *URLHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	if !h.checkMethod(w, r, http.MethodGet) {
		return
	}

	activeUser, ok := h.getUserFromToken(w, r)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	balance, err := userservices.GetBalance(h.db, r.Context(), activeUser.ID)
	if err != nil {
		app.Sugar.Errorw(err.Error(), "event", "get balance")
		h.StatusServerError(w, r)
		return
	}

	resp, jsonErr := json.Marshal(balance)

	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
