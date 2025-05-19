package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kirsh-nat/gophermart.git/internal/models/user"
)

var dataUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

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

func (h *URLHandler) setCookieToken(user *user.User, w http.ResponseWriter) (*user.User, bool) {
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

func (h *URLHandler) getUserFromToken(w http.ResponseWriter, r *http.Request) (*user.User, bool) {
	cookieToken, err := r.Cookie("token")
	if err != nil || cookieToken == nil {
		return &user.User{}, false
	}

	userId, err := getUserID(cookieToken.Value)
	if err != nil {
		return &user.User{}, false
	}

	userModel := user.NewUserModel(h.db)
	foundUser, err := userModel.GetById(r.Context(), userId)
	if err != nil {
		return &user.User{}, false
	}

	user := foundUser.(*user.User)

	return user, true
}

func getUserID(tokenString string) (int, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(secretKey), nil
		})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, err
	}

	return claims.UserID, nil
}
