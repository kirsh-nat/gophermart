package draft

import (
	"context"
	"errors"
)

func (draftModel *DraftModel) Create(ctx context.Context, anyModel any) (any, error) {
	model, ok := (anyModel).(*Draft)
	if !ok {
		return &Draft{}, errors.New("invalid model type")
	}

	_, err := draftModel.DB.ExecContext(ctx,
		"INSERT INTO drafts (number, user_id, sum) VALUES ($1, $2, $3)", model.Number, model.userID, model.Sum)

	if err != nil {
		return model, err
	}

	return model, nil
}
