package user

import (
	"context"
	"errors"
)

func (userModel *UserModel) Update(ctx context.Context, model any) (any, error) {
	//func (userModel *UserModel) Delete(ctx context.Context, model *any) error {
	user, ok := (model).(*User)
	if !ok {
		return &User{}, errors.New("invalid model type")
	}

	_, err := userModel.db.ExecContext(ctx,
		"INSERT INTO users (id, username, password, balance, spent, created_at) VALUES ($1, $2)",
		user.ID, user.Login, user.Password, user.Balance, user.Spent, user.CreatedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}
