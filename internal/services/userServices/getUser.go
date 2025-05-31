package userservices

import (
	"context"
	"database/sql"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

func FindOne(DB *sql.DB, ctx context.Context, login, password string) (*models.User, error) {
	user := &models.User{}
	err := DB.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE username = $1", login).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewAuthorizationError("Cannot find user", err)
		}
		return nil, err
	}
	authorizationError := checkPassword(user.Password, password)
	if authorizationError != nil {
		return nil, NewAuthorizationError("Cannot find user", authorizationError)
	}

	return user, nil
}
