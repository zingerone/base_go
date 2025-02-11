// @Author: Imam Magribi
// @Package: db
package db

import (
	"context"
	"gorm.io/gorm"
)

// CommitOrRollback commits or rolls back a transaction
func (d *DB) CommitOrRollback(ctx context.Context, err error) error {
	if d.db == nil {
		return nil
	}
	tx, ok := ctx.Value(DBTX).(*gorm.DB)
	if !ok {
		return nil
	}
	if err != nil {
		return tx.Rollback().Error

	}
	return tx.Commit().Error
}
