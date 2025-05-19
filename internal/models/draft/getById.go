package draft

import (
	"context"
	"database/sql"
)

func (draftModel *DraftModel) GetByID(ctx context.Context, id int) (any, error) {
	var draft Draft
	err := draftModel.DB.QueryRowContext(ctx,
		"SELECT id, user_id, number, sum, processed_at FROM drafts WHERE id = $1", id).Scan(&draft.ID, &draft.UserID, &draft.Number, &draft.Sum, &draft.ProcessedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &Draft{}, nil
		}
		return &Draft{}, err
	}

	return &draft, nil
}
