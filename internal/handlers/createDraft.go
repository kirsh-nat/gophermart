package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kirsh-nat/gophermart.git/gophermart/internal/models/draft"
	"github.com/kirsh-nat/gophermart.git/gophermart/internal/models/order"
	"github.com/kirsh-nat/gophermart.git/gophermart/internal/models/user"
)

type DraftItem struct {
	Number string  `json:"order"`
	Sum    float32 `json:"sum"`
}

func (h *URLHandler) CreateDraft(w http.ResponseWriter, r *http.Request) {
	if !h.checkMethod(w, r, http.MethodPost) {
		return
	}

	activeUser, ok := h.getUserFromToken(w, r)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	var draftItem DraftItem

	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &draftItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderModel := order.OrderModel{DB: h.db}
	activeOrder, err := orderModel.GetByNumber(r.Context(), draftItem.Number)
	if err != nil {
		var notFoundErr *order.OrderNotFoundError
		if errors.As(err, &notFoundErr) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		h.StatusServerError(w, r)
		return
	}

	if activeOrder.UserID != activeUser.ID {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if activeOrder.Accural < draftItem.Sum {
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	draftModel := draft.DraftModel{DB: h.db}
	newDraft := &draft.Draft{Number: draftItem.Number, UserID: activeUser.ID, Sum: draftItem.Sum}
	_, err = draftModel.Create(r.Context(), newDraft)
	if err != nil {
		h.StatusServerError(w, r)
		return
	}

	userModel := user.UserModel{DB: h.db}
	err = userModel.UpdateBalance(r.Context(), activeUser.ID, draftItem.Sum)
	if err != nil {
		fmt.Println(err, "ERrOR")

		h.StatusServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Payment success"))
}
