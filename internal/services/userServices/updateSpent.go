package userservices

import (
	"context"
	"database/sql"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

func UpdateSpent(DB *sql.DB, ctx context.Context, userID int, newBalance, newSpent float32) error {
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var user models.User
	err = tx.QueryRowContext(ctx, "SELECT balance, spent FROM users WHERE id = $1 FOR UPDATE", userID).Scan(&user.Balance, &user.Spent)
	if err != nil {
		if err == sql.ErrNoRows {
			return NewUserNotFoundError("User not found", err)
		}
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE users SET balance = $1, spent = $2 WHERE id = $3", newBalance, newSpent, userID)
	if err != nil {
		return err
	}

	return nil
}
