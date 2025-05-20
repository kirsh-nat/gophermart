package order

import (
	"context"
	"database/sql"
)

func GetActiveOrders(sqlDB *sql.DB, ctx context.Context) ([]string, error) {
	var orderNums []string
	rows, err := sqlDB.QueryContext(ctx, "SELECT number FROM orders WHERE status in ($1, $2, $3)", REGISTERED, PROCESSING, NEW)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var number string
		if err := rows.Scan(&number); err != nil {
			return orderNums, err
		}
		orderNums = append(orderNums, number)
	}

	if err := rows.Err(); err != nil {
		return orderNums, err
	}

	return orderNums, nil
}
