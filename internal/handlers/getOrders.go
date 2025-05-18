package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kirsh-nat/gophermart.git/gophermart/internal/models/order"
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

	orderModel := &order.OrderModel{DB: h.db}
	orders, err := orderModel.GetUserList(r.Context(), user.ID)
	if err != nil {
		fmt.Println(err, "GetOrders ERROR")
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

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
