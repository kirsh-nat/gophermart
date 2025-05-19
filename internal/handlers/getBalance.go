package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kirsh-nat/gophermart.git/gophermart/internal/app"
	"github.com/kirsh-nat/gophermart.git/gophermart/internal/models/user"
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

	userModel := &user.UserModel{DB: h.db}
	balance, err := userModel.GetBalance(r.Context(), activeUser.ID)
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
