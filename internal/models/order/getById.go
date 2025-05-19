package order

import (
	"context"
	"database/sql"
)

func (orderModel *OrderModel) GetByID(ctx context.Context, id int) (any, error) {
	var order Order
	err := orderModel.DB.QueryRowContext(ctx,
		"SELECT id, user_id, number, status, updated_at FROM orders WHERE user_id = $1", id).Scan(&order.ID, &order.UserID, &order.Number, &order.Status, &order.UreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &Order{}, nil
		}
		return &Order{}, err
	}

	return &order, nil
}
