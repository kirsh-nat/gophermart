package draftservices

import (
	"context"
	"database/sql"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

func GetByID(DB *sql.DB, ctx context.Context, id int) (*models.Draft, error) {
	var draft models.Draft
	err := DB.QueryRowContext(ctx,
		"SELECT id, user_id, number, sum, processed_at FROM drafts WHERE id = $1", id).Scan(&draft.ID, &draft.UserID, &draft.Number, &draft.Sum, &draft.ProcessedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &draft, nil
}
