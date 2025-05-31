package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddDraftTable, downAddDraftTable)
}

func upAddDraftTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "CREATE TABLE drafts (id SERIAL PRIMARY KEY, number VARCHAR(254) NOT NULL, user_id INT NOT NULL, sum DECIMAL(10, 2) DEFAULT 0.00,  processed_at TIMESTAMP WITH TIME ZONE DEFAULT NOW())")
	if err != nil {
		return err
	}
	return nil
}

func downAddDraftTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE drafts")
	if err != nil {
		return err
	}
	return nil
}
