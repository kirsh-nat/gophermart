package user

import (
	"context"
	"database/sql"
)

func (conf *UserModel) GetById(ctx context.Context, id int) (any, error) {
	var user User
	err := conf.db.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return &User{}, nil
		}
		return &User{}, err
	}
	return &user, nil
}

// result, err := userModel.GetById(ctx, 123)
// if err != nil {
//     // обработка ошибки
// }

// if u, ok := result.(*User); ok {
//     // используйте u как *User
// }
