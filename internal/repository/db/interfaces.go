// @Author: Imam Magribi
// @Package: db
package db

import (
	"context"
	"database/sql"
)

// DBInterface is the interface for the database
type DBInterface interface {
	GetDb() *sql.DB
	BeginTx(ctx context.Context) (context.Context, error)
	CommitOrRollback(ctx context.Context, err error) error
}
