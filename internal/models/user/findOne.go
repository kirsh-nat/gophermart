package user

import (
	"context"
	"database/sql"
)

func (conf *UserModel) FindOne(ctx context.Context, login, password string) (*User, error) {
	var user User
	err := conf.DB.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE username = $1", login).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return &User{}, nil
		}
		return &User{}, err
	}
	authorizationError := checkPassword(user.Password, password)
	if authorizationError != nil {
		return &User{}, NewAuthorizationError("Cannot find user", authorizationError)
	}

	return &user, nil
}
