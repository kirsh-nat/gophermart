package draft

import (
	"database/sql"
	"time"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

type Draft struct {
	ID          int       `json:"id"`           // уникальный идентификатор
	UserID      int       `json:"user_id"`      // имя пользователя
	Number      string    `json:"number"`       // хэшированный пароль
	Sum         float32   `json:"sum"`          // потраченная сумма балллов за весь период регистрации
	ProcessedAt time.Time `json:"processed_at"` // дата создания
}

type DraftModel struct {
	DB *sql.DB
}

func NewDraftModel(db *sql.DB) models.ModelInterface {
	return &DraftModel{DB: db}
}
