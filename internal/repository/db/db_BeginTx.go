// @Author: Imam Magribi
// @Package: db
package db

import (
	"context"
	"database/sql"
)

// BeginTx begins a transaction
func (d *DB) BeginTx(ctx context.Context, txOptions *TxOptions) (context.Context, error) {
	// BeginTx begins a transaction
	if ctx == nil {
		ctx = context.Background()
	}

	ok, _ := ctx.Value(DBTX).(*sql.Tx)
	if ok != nil {
		return ctx, nil
	}
	var exTxOptions *sql.TxOptions
	if txOptions != nil {
		exTxOptions = &sql.TxOptions{
			Isolation: txOptions.IsolationLevel,
			ReadOnly:  txOptions.ReadOnly,
		}
	}

	tx, err := d.db.BeginTx(ctx, exTxOptions)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, DBTX, tx)
	return ctx, nil
}
