package userservices

import (
	"context"
	"database/sql"
)

type Invoice struct {
	Balance float32 `json:"current"`
	Spent   float32 `json:"withdrawn"`
}

func GetBalance(DB *sql.DB, ctx context.Context, id int) (Invoice, error) {
	var invoice Invoice
	err := DB.QueryRowContext(ctx, "SELECT balance, spent FROM users WHERE id = $1", id).Scan(&invoice.Balance, &invoice.Spent)
	if err != nil {
		if err == sql.ErrNoRows {
			return Invoice{}, nil
		}
		return Invoice{}, err
	}

	return invoice, nil
}
