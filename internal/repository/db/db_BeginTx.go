// @Author: Imam Magribi
// @Package: db
package db

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

// BeginTx begins a transaction
func (d *DB) BeginTx(ctx context.Context, txOptions *TxOptions) (context.Context, error) {
	// BeginTx begins a transaction
	if ctx == nil {
		ctx = context.Background()
	}

	_, ok := ctx.Value(DBTX).(*gorm.DB)
	if ok {
		return ctx, nil
	}
	var exTxOptions *sql.TxOptions
	if txOptions != nil {
		exTxOptions = &sql.TxOptions{
			Isolation: txOptions.IsolationLevel,
			ReadOnly:  txOptions.ReadOnly,
		}
	}

	tx := d.db.Begin(exTxOptions)

	ctx = context.WithValue(ctx, DBTX, tx)
	return ctx, nil
}
