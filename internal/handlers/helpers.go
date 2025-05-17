package handlers

import (
	"database/sql"
	"net/http"

	"github.com/kirsh-nat/gophermart.git/gophermart/internal/models/user"
)

// type URLHandler struct {
// 	service models.Model
// }

// func NewURLHandler(service *models.Model) *URLHandler {
// 	return &URLHandler{service: *service}
// }

type URLHandler struct {
	db *sql.DB
	// service models.Model
}

// NewURLHandler returns a new instance of URLHandler with the given service.
//
// The service field of the returned URLHandler is a copy of the given service.
func NewURLHandler(db *sql.DB) *URLHandler {
	return &URLHandler{db: db}
}

// func (h *URLHandler) setCookieToken(w http.ResponseWriter, r *http.Request) (*user.User, bool) {
// 	cookieToken, err := r.Cookie("token")
// 	if err != nil || cookieToken == nil {
// 		return h.createUserAndSetCookie(w, r)
// 	}

// 	user, err := user.GetUser(cookieToken.Value)
// 	if err != nil {
// 		return h.createUserAndSetCookie(w)
// 	}

// 	return user, true
// }

func (h *URLHandler) setCookieToken(user *user.User, w http.ResponseWriter, r *http.Request) (*user.User, bool) {
	// userModel := user.NewUserModel(h.service.DB)
	// u, err := userModel.Create(r.Context(), &user.User{})
	// if err != nil {
	// 	return nil, false
	// }
	// user, ok := (u).(*user.User)
	// if !ok {
	// 	return nil, false
	// }

	token, err := createToken(user)
	if err != nil {
		return nil, false
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
		Path:  "/",
	})
	return user, true
}
