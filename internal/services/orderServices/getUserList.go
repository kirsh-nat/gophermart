package orderservices

import (
	"context"
	"database/sql"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

func GetUserList(DB *sql.DB, ctx context.Context, userID int) ([]models.OrderItem, error) {
	var orders []models.OrderItem
	rows, err := DB.QueryContext(ctx,
		"SELECT  number, status, accural, updated_at FROM orders WHERE user_id = $1 ORDER BY updated_at desc", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return orders, nil
		}
		return orders, err
	}

	defer rows.Close()

	for rows.Next() {
		var order models.OrderItem
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
