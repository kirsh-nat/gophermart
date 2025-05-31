package userservices

import (
	"context"
	"database/sql"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

func UpdateSpent(DB *sql.DB, ctx context.Context, userID int, sum float32) error {
	var user models.User
	err := DB.QueryRowContext(ctx, "SELECT balance, spent FROM users WHERE id = $1 ", userID).Scan(&user.Balance, &user.Spent)
	if err != nil {
		if err == sql.ErrNoRows {
			return NewUserNotFoundError("User not found", err)
		}
		return err
	}

	newBalance := user.Balance - sum
	newSpent := user.Spent + sum

	_, err = DB.ExecContext(ctx, "UPDATE users SET balance = $1, spent = $2 WHERE id = $3", newBalance, newSpent, userID)
	if err != nil {
		return err
	}

	return nil
}
