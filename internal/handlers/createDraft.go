package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	draftservices "github.com/kirsh-nat/gophermart.git/internal/services/draftServices"
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

	err = draftservices.NewUserDraft(h.db, r.Context(), activeUser, draftItem.Number, draftItem.Sum)
	if err != nil {
		var userNotAuthErr *draftservices.UserNotAuthorizedError
		if errors.As(err, &userNotAuthErr) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		var paymentRequiredErr *draftservices.PaymentRequiredError
		if errors.As(err, &paymentRequiredErr) {
			w.WriteHeader(http.StatusPaymentRequired)
			return
		}

		h.StatusServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Payment success"))
}
