// @Author: Imam Magribi
// @Package: db
package db

import (
	"context"
	"gorm.io/gorm"
)

// GetDb returns the database connection
func (d *DB) GetDb(ctx context.Context) *gorm.DB {
	val, ok := ctx.Value(DBTX).(*gorm.DB)
	if ok {
		return val
	}
	return d.db
}
