package draftservices

import (
	"context"
	"database/sql"

	"github.com/kirsh-nat/gophermart.git/internal/models"
)

func CreateDraft(DB *sql.DB, ctx context.Context, number string, userID int, sum float32) (*models.Draft, error) {
	var dr models.Draft
	_, err := DB.ExecContext(ctx,
		"INSERT INTO drafts (number, user_id, sum) VALUES ($1, $2, $3)", number, userID, sum)

	if err != nil {
		return nil, err
	}

	return &dr, nil
}
