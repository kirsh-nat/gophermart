package orderservices

import (
	"context"
	"database/sql"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

func GetByID(DB *sql.DB, ctx context.Context, id int) (*models.Order, error) {
	var order models.Order
	err := DB.QueryRowContext(ctx,
		"SELECT id, user_id, number, status, updated_at FROM orders WHERE user_id = $1", id).Scan(&order.ID, &order.UserID, &order.Number, &order.Status, &order.UreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}
