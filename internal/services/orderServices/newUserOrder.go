package orderservices

import (
	"context"
	"database/sql"
)

func NewUserOrder(DB *sql.DB, ctx context.Context, userID int, number string) error {
	err := CheckNumber(number)
	if err != nil {
		return err
	}

	_, err = Create(DB, ctx, number, userID)
	if err != nil {
		return err
	}

	return nil
}
