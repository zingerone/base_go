// @Author: Imam Magribi
// @Package: db
package db

import (
	"context"
	"gorm.io/gorm"
)

// DBInterface is the interface for the database
type DBInterface interface {
	GetDb(ctx context.Context) *gorm.DB
	BeginTx(ctx context.Context) (context.Context, error)
	CommitOrRollback(ctx context.Context, err error) error
}
