package user

import (
	"context"
	"errors"
)

func (userModel *UserModel) Delete(ctx context.Context, model any) error {
	user, ok := (model).(*User)
	if !ok {
		return errors.New("invalid model type")
	}

	_, err := userModel.db.ExecContext(ctx,
		"DELETE FROM users WHERE id = ($1)",
		user.ID)

	if err != nil {
		return err
	}

	return nil
}
