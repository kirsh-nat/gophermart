package models

import "time"

// модель заказа
type Order struct {
	ID        int       `json:"id"`         // уникальный идентификатор
	UserID    int       `json:"userID"`     // имя пользователя
	Number    string    `json:"number"`     // хэшированный пароль
	Accural   float32   `json:"accural"`    // потраченная сумма балллов за весь период регистрации
	Status    float32   `json:"status"`     // текущий баланс
	UreatedAt time.Time `json:"updated_at"` // дата создания
}

type OrderItem struct {
	Number    string    `json:"number"`
	Accural   float32   `json:"accural"`
	Status    string    `json:"status"`
	UreatedAt time.Time `json:"updated_at"`
}

// модель пользователя

type User struct {
	ID        int       `json:"id"`         // уникальный идентификатор
	Login     string    `json:"login"`      // имя пользователя
	Password  string    `json:"password"`   // хэшированный пароль
	Spent     float32   `json:"spent"`      // потраченная сумма балллов за весь период регистрации
	Balance   float32   `json:"balance"`    // текущий баланс
	CreatedAt time.Time `json:"created_at"` // дата создания
}

// модель драфта
type Draft struct {
	ID          int       `json:"id"`           // уникальный идентификатор
	UserID      int       `json:"user_id"`      // имя пользователя
	Number      string    `json:"number"`       // хэшированный пароль
	Sum         float32   `json:"sum"`          // потраченная сумма балллов за весь период регистрации
	ProcessedAt time.Time `json:"processed_at"` // дата создания
}

type DraftItem struct {
	Order       string    `json:"order"`
	Sum         float32   `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}
