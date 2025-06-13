package orderservices

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kirsh-nat/gophermart.git/internal/models"
)

func Create(DB *sql.DB, ctx context.Context, number string, userID int) (*models.Order, error) {

	checkOrder, err := GetByNumber(DB, ctx, number)
	if err != nil {
		notFoundErr := &OrderNotFoundError{"Order not found", err}
		if errors.As(err, &notFoundErr) {
		} else {
			return nil, err
		}
	} else {
		if checkOrder.UserID == userID {
			return nil, NewUserNumberExistsError("Create Order", err)
		}
		return nil, NewNumberExists("Create Order", err)
	}

	result, err := DB.ExecContext(ctx,
		"INSERT INTO Orders (number, user_id) VALUES ($1, $2)", number, userID)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return nil, NewNumberExists("Create Order", err)
		}
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, errors.New("failed to create order")
	}

	return &models.Order{Number: number, UserID: userID}, nil
}
