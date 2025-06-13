package handlers

import (
	"encoding/json"
	"net/http"

	orderservices "github.com/kirsh-nat/gophermart.git/internal/services/orderServices"
)

func (h *URLHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	if !h.checkMethod(w, r, http.MethodGet) {
		return
	}

	user, ok := h.getUserFromToken(w, r)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	orders, err := orderservices.GetUserList(h.db, r.Context(), user.ID)
	if err != nil {
		h.StatusServerError(w, r)
		return
	}

	if len(orders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	resp, jsonErr := json.Marshal(orders)

	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
