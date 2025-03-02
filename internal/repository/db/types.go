// @Author: Imam Magribi
// @Package: db
package db

import (
	"database/sql"
	"gorm.io/gorm"
)

// DB is an interface for database operations
type DriverName string

// NewDB creates a new database connection
func (d DriverName) String() string {
	return string(d)
}

// DB is an interface for database operations
type DB struct {
	db *gorm.DB
}

type IsolationLevel = sql.IsolationLevel
type TxOptions struct {
	IsolationLevel IsolationLevel
	ReadOnly       bool
}
