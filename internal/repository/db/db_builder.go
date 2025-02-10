// @Author: Imam Magribi
// @Package: db
package db

import (
	"database/sql"
	"github.com/zingerone/base_go/utils/pointer"
	"time"
)

// New creates a new database connection
func New(driverName DriverName, dataSource string) *DBBuilder {
	return &DBBuilder{
		config:     &DbConfig{},
		driverName: driverName,
		dataSource: dataSource,
	}
}

type DBBuilder struct {
	driverName DriverName
	dataSource string
	config     *DbConfig
}

// WithMaxLifeTime this is for max life time
func (b *DBBuilder) WithMaxLifeTime(duration time.Duration) *DBBuilder {
	b.config.maxIdleLifeTime = &duration
	return b
}

// WithMaxIdleConnections this is for max idle connections
func (b *DBBuilder) WithMaxIdleConnections(maxIdle int) *DBBuilder {
	b.config.maxIdleConnection = &maxIdle
	return b
}

// WithMaxOpenConn this is for max open connections
func (b *DBBuilder) WithMaxOpenConn(maxOpenConns int) *DBBuilder {
	b.config.maxOpenConnection = &maxOpenConns
	return b
}

// Open opens a new database connection
func (b *DBBuilder) Open() (*DB, error) {
	db, err := sql.Open(b.driverName.String(), b.dataSource)
	if err != nil {
		return nil, err
	}
	if b.config != nil {
		pointer.SetIfNotNil(b.config.maxIdleLifeTime, func(value *time.Duration) {
			db.SetConnMaxLifetime(*value)
		})
		pointer.SetIfNotNil(b.config.maxIdleConnection, func(value *int) {
			db.SetMaxIdleConns(*value)
		})

		pointer.SetIfNotNil(b.config.maxOpenConnection, func(value *int) {
			db.SetMaxOpenConns(*value)
		})
	}
	return &DB{db: db}, nil
}
