package draftservices

import (
	"context"
	"database/sql"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

func GetUserList(DB *sql.DB, ctx context.Context, userID int) ([]models.DraftItem, error) {
	var drafts []models.DraftItem
	rows, err := DB.QueryContext(ctx,
		"SELECT number, sum, processed_at FROM drafts WHERE user_id = $1 ORDER BY processed_at desc", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var draft models.DraftItem
		if err := rows.Scan(&draft.Order, &draft.Sum, &draft.ProcessedAt); err != nil {
			return nil, err
		}
		drafts = append(drafts, draft)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return drafts, nil
}
