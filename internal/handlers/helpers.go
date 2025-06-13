package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kirsh-nat/gophermart.git/internal/models"
	userservices "github.com/kirsh-nat/gophermart.git/internal/services/userServices"
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
	UserID int
}

func NewURLHandler(db *sql.DB) *URLHandler {
	return &URLHandler{db: db}
}

func (h *URLHandler) setCookieToken(user *models.User, w http.ResponseWriter) (*models.User, bool) {
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

func (h *URLHandler) getUserFromToken(w http.ResponseWriter, r *http.Request) (*models.User, bool) {
	cookieToken, err := r.Cookie("token")
	if err != nil || cookieToken == nil {
		return &models.User{}, false
	}

	userID, err := getuserID(cookieToken.Value)
	if err != nil {
		return &models.User{ID: userID}, false
	}

	//userModel := user.NewUserModel(h.db)
	foundUser, err := userservices.GetByID(h.db, r.Context(), userID)
	if err != nil {
		return &models.User{}, false
	}

	//found := foundUser.(*models.User)

	return foundUser, true
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
func createToken(user *models.User) (string, error) {
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
