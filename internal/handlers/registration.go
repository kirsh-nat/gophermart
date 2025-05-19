package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kirsh-nat/gophermart.git/internal/models/user"
)

func (h *URLHandler) Registration(w http.ResponseWriter, r *http.Request) {
	if !h.checkMethod(w, r, http.MethodPost) {
		return
	}

	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		h.StatusBadRequest(w, r)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &dataUser); err != nil {
		h.StatusBadRequest(w, r)
		return
	}

	userModel := user.NewUserModel(h.db)
	u, err := userModel.Create(r.Context(), &user.User{Login: dataUser.Login, Password: dataUser.Password})
	if err != nil {
		var dErr *user.UserExistsError
		if errors.As(err, &dErr) {
			w.WriteHeader(http.StatusConflict)
			return

		}
		h.StatusServerError(w, r)
		return
	}

	user, ok := (u).(*user.User)
	if !ok {
		h.StatusServerError(w, r)
		return
	}

	user, ok = h.setCookieToken(user, w)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something went wrong"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("User was successfully created: " + user.Login))
}
