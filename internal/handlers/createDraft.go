package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	draftservices "github.com/kirsh-nat/gophermart.git/internal/services/draftServices"
	orderservices "github.com/kirsh-nat/gophermart.git/internal/services/orderServices"
	userservices "github.com/kirsh-nat/gophermart.git/internal/services/userServices"
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

	activeOrder, err := orderservices.GetByNumber(h.db, r.Context(), draftItem.Number)
	if err != nil {
		var notFoundErr *orderservices.OrderNotFoundError
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

	_, err = draftservices.CreateDraft(h.db, r.Context(), draftItem.Number, activeUser.ID, draftItem.Sum)
	if err != nil {
		h.StatusServerError(w, r)
		return
	}

	err = userservices.UpdateSpent(h.db, r.Context(), activeUser.ID, draftItem.Sum)
	if err != nil {
		h.StatusServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Payment success"))
}
