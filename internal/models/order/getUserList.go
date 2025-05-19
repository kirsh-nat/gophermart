package order

import (
	"context"
	"database/sql"
	"time"
)

type OrderItem struct {
	Number    string    `json:"number"`
	Accural   float32   `json:"accural"` // TODO: только если существует
	Status    string    `json:"status"`
	UreatedAt time.Time `json:"updated_at"`
}

func (orderModel *OrderModel) GetUserList(ctx context.Context, userID int) ([]OrderItem, error) {
	var orders []OrderItem
	rows, err := orderModel.DB.QueryContext(ctx,
		"SELECT  number, status, accural, updated_at FROM orders WHERE user_id = $1 ORDER BY updated_at desc", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return orders, nil
		}
		return orders, err
	}

	defer rows.Close()

	for rows.Next() {
		var order OrderItem
		if err := rows.Scan(&order.Number, &order.Status, &order.Accural, &order.UreatedAt); err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return orders, err
	}

	return orders, nil
}
