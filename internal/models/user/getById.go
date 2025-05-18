package user

import (
	"context"
	"database/sql"
)

func (conf *UserModel) GetById(ctx context.Context, id int) (any, error) {
	var user User
	err := conf.DB.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return &User{}, nil
		}
		return &User{}, err
	}
	return &user, nil
}
