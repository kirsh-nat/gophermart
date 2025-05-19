package order

import (
	"context"
	"database/sql"
)

func (orderModel OrderModel) GetByNumber(ctx context.Context, number string) (*Order, error) {
	var order Order
	err := orderModel.DB.QueryRowContext(ctx, "SELECT id, user_id, number, accural FROM orders WHERE number = $1", number).Scan(&order.ID, &order.userID, &order.Number, &order.Accural)
	if err != nil {
		if err == sql.ErrNoRows {
			return &Order{}, NewOrderNotFoundError("Order not found", err)
		}
		return &Order{}, err
	}

	return &order, nil
}
