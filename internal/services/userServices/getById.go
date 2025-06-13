package userservices

import (
	"context"
	"database/sql"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

func GetByID(DB *sql.DB, ctx context.Context, id int) (*models.User, error) {
	user := &models.User{}
	err := DB.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}
	return user, nil
}
