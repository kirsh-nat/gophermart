package orderservices

import (
	"context"
	"database/sql"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

func GetByNumber(DB *sql.DB, ctx context.Context, number string) (*models.Order, error) {
	order := &models.Order{}
	err := DB.QueryRowContext(ctx, "SELECT id, user_id, number, accural FROM orders WHERE number = $1", number).Scan(&order.ID, &order.UserID, &order.Number, &order.Accural)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewOrderNotFoundError("Order not found", err)
		}
		return nil, err
	}

	return order, nil
}
