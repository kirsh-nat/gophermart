package draft

import (
	"context"
	"errors"
)

func (draftModel *DraftModel) Update(ctx context.Context, model any) (any, error) {
	draft, ok := (model).(*Draft)
	if !ok {
		return &Draft{}, errors.New("invalid model type")
	}

	return draft, nil
}
