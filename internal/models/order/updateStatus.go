package order

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

const (
	PROCESSED  = "PROCESSED"
	REGISTERED = "REGISTERED"
	INVALID    = "INVALID"
	PROCESSING = "PROCESSING"
	NEW        = "NEW"
)

func UpdateStatus(sqlDb *sql.DB, ctx context.Context, result Result) error {

	if (result.Status != PROCESSED) && (result.Status != REGISTERED) && (result.Status != INVALID) && (result.Status != PROCESSING) {
		return errors.New("invalid status")
	}

	fmt.Println("result.Status", result.Status, result.Order, result.Accrual)

	_, err := sqlDb.ExecContext(ctx, "UPDATE orders SET status = $1, accural = $3 WHERE number = $2", result.Status, result.Order, result.Accrual)
	if err != nil {
		return err
	}
	return nil
}
