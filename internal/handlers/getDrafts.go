package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kirsh-nat/gophermart.git/gophermart/internal/app"
	"github.com/kirsh-nat/gophermart.git/gophermart/internal/models/draft"
)

func (h *URLHandler) GetDrafts(w http.ResponseWriter, r *http.Request) {
	if !h.checkMethod(w, r, http.MethodGet) {
		return
	}

	user, ok := h.getUserFromToken(w, r)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	draftModel := &draft.DraftModel{DB: h.db}
	drafts, err := draftModel.GetUserList(r.Context(), user.ID)
	if err != nil {
		app.Sugar.Errorw(err.Error(), "event", "get drafts")
		h.StatusServerError(w, r)
		return
	}

	if len(drafts) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	resp, jsonErr := json.Marshal(drafts)

	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
