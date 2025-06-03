package userservices

import (
	"context"
	"database/sql"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

func UpdateBalance(sqlDB *sql.DB, ctx context.Context, userID int, balance float32) error {
	var user models.User
	err := sqlDB.QueryRowContext(ctx, "SELECT balance, spent FROM users WHERE id = $1 ", userID).Scan(&user.Balance, &user.Spent)
	if err != nil {
		if err == sql.ErrNoRows {
			return NewUserNotFoundError("User not found", err)
		}
		return err
	}

	newBalance := user.Balance + balance

	_, err = sqlDB.ExecContext(ctx, "UPDATE users SET balance = $1  WHERE id = $2", newBalance, userID)
	if err != nil {
		return err
	}

	return nil
}
