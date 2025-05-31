package orderservices

import (
	"context"
	"database/sql"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

func Withdrawn(DB *sql.DB, ctx context.Context, number string, userID int, sum float32) error {
	var order models.Order
	err := DB.QueryRowContext(ctx, "SELECT status, accural FROM orders WHERE number = $1 AND user_id = $2", number, userID).Scan(&order.ID, &order.UserID, &order.Number)
	if err != nil {
		if err == sql.ErrNoRows {
			return NewOrderNotFoundError("Order not found", err)
		}
		return err
	}

	if err = checkBalance(sum, order.Accural); err != nil {
		return err
	}

	newSum := order.Accural - sum
	_, err = DB.ExecContext(ctx, "UPDATE orders SET accural = $1 WHERE number = $2", newSum, number)
	if err != nil {
		return err
	}

	return nil
}
