package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kirsh-nat/gophermart.git/internal/models/user"
)

var dataUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type URLHandler struct {
	db *sql.DB
}

const tokenExp = time.Hour * 3
const secretKey = "supersecretkey"

type Claims struct {
	jwt.RegisteredClaims
	UserID int `json:"user_id"`
}

func NewURLHandler(db *sql.DB) *URLHandler {
	return &URLHandler{db: db}
}

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
	fmt.Println("set cookieToken: ", token)
	return user, true
}

func (h *URLHandler) getUserFromToken(w http.ResponseWriter, r *http.Request) (*user.User, bool) {
	cookieToken, err := r.Cookie("token")
	fmt.Println("get cookieToken: ", cookieToken)
	if err != nil || cookieToken == nil {
		return &user.User{}, false
	}

	userID, err := getuserID(cookieToken.Value)
	if err != nil {
		return &user.User{}, false
	}

	userModel := user.NewUserModel(h.db)
	foundUser, err := userModel.GetByID(r.Context(), userID)
	if err != nil {
		return &user.User{}, false
	}

	found := foundUser.(*user.User)

	return found, true
}

func getuserID(tokenString string) (int, error) {
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
		return 0, fmt.Errorf("invalid token")
	}

	return claims.UserID, nil
}
func createToken(user *user.User) (string, error) {
	token, err := buildJWTString(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}

func buildJWTString(UUID int) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExp)),
		},
		UserID: UUID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
