package order

import (
	"context"
	"database/sql"
)

func (orderModel *OrderModel) getByNumber(ctx context.Context, number string) (*Order, error) {
	var order Order
	err := orderModel.DB.QueryRowContext(ctx, "SELECT id, user_id, number FROM orders WHERE number = $1", number).Scan(&order.ID, &order.UserID, &order.Number)
	if err != nil {
		if err == sql.ErrNoRows {
			return &Order{}, nil
		}
		return &Order{}, err
	}

	return &order, nil
}
