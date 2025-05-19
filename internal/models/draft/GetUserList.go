package draft

import (
	"context"
	"time"
)

type DraftItem struct {
	Order       string    `json:"order"`
	Sum         float32   `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}

func (conf *DraftModel) GetUserList(ctx context.Context, userID int) ([]DraftItem, error) {
	var drafts []DraftItem
	rows, err := conf.DB.QueryContext(ctx,
		"SELECT number, sum, processed_at FROM drafts WHERE user_id = $1 ORDER BY processed_at desc", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var draft DraftItem
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
