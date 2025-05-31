package systemservices

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kirsh-nat/gophermart.git/internal/models/order"
)

func UpdateStatus(sqlDB *sql.DB, ctx context.Context, result Result) (int, error) {
	if result.Status != order.PROCESSED && result.Status != order.REGISTERED && result.Status != order.INVALID && result.Status != order.PROCESSING {
		return 0, errors.New("invalid status")
	}

	_, err := sqlDB.ExecContext(ctx,
		"UPDATE orders SET status = $1, accural = $3 WHERE number = $2",
		result.Status,
		result.Order,
		result.Accrual,
	)
	if err != nil {
		return 0, err
	}

	var userID int
	err = sqlDB.QueryRowContext(ctx,
		"SELECT user_id FROM orders WHERE number = $1",
		result.Order,
	).Scan(&userID)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("user not found")
		}
		return 0, err
	}

	return userID, nil
}

// func UpdateStatus(sqlDB *sql.DB, ctx context.Context, result Result) int, error {

// 	if (result.Status != PROCESSED) && (result.Status != REGISTERED) && (result.Status != INVALID) && (result.Status != PROCESSING) {
// 		return 0, errors.New("invalid status")
// 	}

// 	fmt.Println("result.Status", result.Status, result.Order, result.Accrual)

// 	_, err := sqlDB.ExecContext(ctx, "UPDATE orders SET status = $1, accural = $3 WHERE number = $2", result.Status, result.Order, result.Accrual)
// 	if err != nil {
// 		return 0,err
// 	}

// 	var userID int
// 	err = sqlDB.QueryRowContext(ctx, "SELECT user_id FROM orders WHERE number = $1", result.Order).Scan(&userID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			// ничего не найдено
// 		} else {
// 			return 0, err
// 		}
// 	}

// 	return userID, nil
// }
