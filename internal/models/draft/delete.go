package draft

import (
	"context"
	"errors"
)

func (draftModel *DraftModel) Delete(ctx context.Context, model any) error {
	draft, ok := (model).(*Draft)
	if !ok {
		return errors.New("invalid model type")
	}

	_, err := draftModel.DB.ExecContext(ctx,
		"DELETE FROM drafts WHERE id = ($1)",
		draft.ID)

	if err != nil {
		return err
	}

	return nil
}
