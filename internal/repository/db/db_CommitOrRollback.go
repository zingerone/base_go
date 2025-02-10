// @Author: Imam Magribi
// @Package: db
package db

import (
	"context"
	"database/sql"
)

// CommitOrRollback commits or rolls back a transaction
func (d *DB) CommitOrRollback(ctx context.Context, err error) error {
	if d.db == nil {
		return nil
	}
	tx, ok := ctx.Value(DBTX).(*sql.Tx)
	if !ok {
		return nil
	}
	if err != nil {
		return tx.Rollback()

	}
	return tx.Commit()
}
