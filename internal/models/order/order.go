package order

import (
	"database/sql"
	"time"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

type Order struct {
	ID        int       `json:"id"`         // уникальный идентификатор
	userID    int       `json:"userID"`     // имя пользователя
	Number    string    `json:"number"`     // хэшированный пароль
	Accural   float32   `json:"accural"`    // потраченная сумма балллов за весь период регистрации
	Status    float32   `json:"status"`     // текущий баланс
	UreatedAt time.Time `json:"updated_at"` // дата создания
}

type OrderModel struct {
	DB *sql.DB
}

func NewOrderModel(db *sql.DB) models.ModelInterface {
	//return &models.Model{db: db}
	return &OrderModel{DB: db}
}
