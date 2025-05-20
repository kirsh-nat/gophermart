package user

import (
	"database/sql"
	"time"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

type User struct {
	ID        int       `json:"id"`         // уникальный идентификатор
	Login     string    `json:"login"`      // имя пользователя
	Password  string    `json:"password"`   // хэшированный пароль
	Spent     float32   `json:"spent"`      // потраченная сумма балллов за весь период регистрации
	Balance   float32   `json:"balance"`    // текущий баланс
	CreatedAt time.Time `json:"created_at"` // дата создания
}

type UserModel struct {
	DB *sql.DB
}

func NewUserModel(db *sql.DB) models.ModelInterface {
	return &UserModel{DB: db}
}
