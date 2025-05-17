package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kirsh-nat/gophermart.git/gophermart/internal/models/user"
)

var dataUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *URLHandler) Registration(w http.ResponseWriter, r *http.Request) {
	if !h.checkMethod(w, r, http.MethodPost) {
		return
	}

	fmt.Println("hello")

	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &dataUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userModel := user.NewUserModel(h.db)
	u, err := userModel.Create(r.Context(), &user.User{Login: dataUser.Login, Password: dataUser.Password})
	if err != nil {
		h.StatusBadRequest(w, r)
		return
	}

	user, ok := (u).(*user.User)
	if !ok {
		h.StatusBadRequest(w, r)
		return
	}
	user, ok = h.setCookieToken(user, w, r)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something went wrong"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("User was successfully created: " + user.Login))
}
