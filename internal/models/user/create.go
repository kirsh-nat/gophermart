package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func (userModel *UserModel) Create(ctx context.Context, model any) (any, error) {
	user, ok := (model).(*User)
	if !ok {
		return &User{}, errors.New("invalid model type")
	}

	hashPassword, err := hashPassword(user.Password)
	if err != nil {
		return user, err
	}

	_, err = userModel.DB.ExecContext(ctx,
		"INSERT INTO users (username, password) VALUES ($1, $2)", user.Login, hashPassword)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return user, NewUserExistsError("Create user", err)
		}
		return user, err
	}

	return user, nil
}
