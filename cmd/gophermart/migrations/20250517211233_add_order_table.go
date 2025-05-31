package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddOrderTable, downAddOrderTable)
}

func upAddOrderTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "CREATE TABLE orders (id SERIAL PRIMARY KEY, number VARCHAR(254) NOT NULL UNIQUE, user_id INT NOT NULL, status VARCHAR(10) NOT NULL DEFAULT 'NEW', accural DECIMAL(10, 2) DEFAULT 0.00)")
	if err != nil {
		return err
	}
	return nil
}

func downAddOrderTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE orders")
	if err != nil {
		return err
	}
	return nil
}
