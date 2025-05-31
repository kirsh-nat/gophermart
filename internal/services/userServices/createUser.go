package userservices

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kirsh-nat/gophermart.git/internal/models"
)

func CreateUser(DB *sql.DB, ctx context.Context, login, password string) (*models.User, error) {
	hashPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	result, err := DB.ExecContext(ctx,
		"INSERT INTO users (username, password) VALUES ($1, $2)", login, hashPassword)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return nil, NewUserExistsError("Create user", err)
		}
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, NewUserExistsError("Create user", err)
	}

	return FindOne(DB, ctx, login, password)
}
