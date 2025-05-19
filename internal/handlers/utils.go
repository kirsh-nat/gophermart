package handlers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kirsh-nat/gophermart.git/internal/models/user"
)

const tokenExp = time.Hour * 3
const secretKey = "supersecretkey"

type Claims struct {
	jwt.RegisteredClaims
	userID int
}

func (h *URLHandler) checkMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return false
	}

	return true
}

func (h *URLHandler) StatusBadRequest(w http.ResponseWriter, r *http.Request) bool {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad request"))
	return false
}

func (h *URLHandler) StatusServerError(w http.ResponseWriter, r *http.Request) bool {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Something went wrong"))
	return false
}

func createToken(user *user.User) (string, error) {
	token, err := buildJWTString(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}

func buildJWTString(UUID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExp)),
		},
		userID: UUID,
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
