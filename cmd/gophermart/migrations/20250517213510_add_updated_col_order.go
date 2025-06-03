package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddUpdatedColOrder, downAddUpdatedColOrder)
}

func upAddUpdatedColOrder(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "ALTER TABLE orders ADD COLUMN updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()")
	if err != nil {
		return err
	}
	return nil
}

func downAddUpdatedColOrder(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "ALTER TABLE orders DROP COLUMN updated_at")
	if err != nil {
		return err
	}
	return nil
}
