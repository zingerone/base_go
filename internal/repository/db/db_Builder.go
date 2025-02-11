// @Author: Imam Magribi
// @Package: db
package db

import (
	"github.com/zingerone/base_go/utils/pointer"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// New creates a new database connection
func New(driverName DriverName, dataSource string) *DBBuilder {
	return &DBBuilder{
		config:     &NewDbConfig{},
		driverName: driverName,
		dsn:        dataSource,
	}
}

type DBBuilder struct {
	driverName DriverName
	dsn        string
	config     *NewDbConfig
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

// WithSkipDefaultTransaction sets the skipDefaultTransaction option
func (b *DBBuilder) WithSkipDefaultTransaction(skipDefaultTransaction bool) *DBBuilder {
	b.config.skipDefaultTransaction = &skipDefaultTransaction
	return b
}

// WithLogger sets the logger option
func (b *DBBuilder) WithLogger(logger logger.Interface) *DBBuilder {
	b.config.logger = &logger
	return b
}

// WithDryRun sets the dryRun option
func (b *DBBuilder) WithDryRun(dryRun bool) *DBBuilder {
	b.config.dryRun = &dryRun
	return b
}

// WithPrepareStmt sets the prepareStmt option
func (b *DBBuilder) WithPrepareStmt(prepareStmt bool) *DBBuilder {
	b.config.prepareStmt = &prepareStmt
	return b
}

// WithDisableAutomaticPing sets the disableAutomaticPing option
func (b *DBBuilder) WithDisableAutomaticPing(disableAutomaticPing bool) *DBBuilder {
	b.config.disableAutomaticPing = &disableAutomaticPing
	return b
}

// WithDisableForeignKeyConstraintWhenMigrating sets the disableForeignKeyConstraintWhenMigrating option
func (b *DBBuilder) WithDisableForeignKeyConstraintWhenMigrating(disableForeignKeyConstraintWhenMigrating bool) *DBBuilder {
	b.config.disableForeignKeyConstraintWhenMigrating = &disableForeignKeyConstraintWhenMigrating
	return b
}

// WithIgnoreRelationshipsWhenMigrating sets the ignoreRelationshipsWhenMigrating option
func (b *DBBuilder) WithIgnoreRelationshipsWhenMigrating(ignoreRelationshipsWhenMigrating bool) *DBBuilder {
	b.config.ignoreRelationshipsWhenMigrating = &ignoreRelationshipsWhenMigrating
	return b
}

// WithDisableNestedTransaction sets the disableNestedTransaction option
func (b *DBBuilder) WithDisableNestedTransaction(disableNestedTransaction bool) *DBBuilder {
	b.config.disableNestedTransaction = &disableNestedTransaction
	return b
}

// WithAllowGlobalUpdate sets the allowGlobalUpdate option
func (b *DBBuilder) WithAllowGlobalUpdate(allowGlobalUpdate bool) *DBBuilder {
	b.config.allowGlobalUpdate = &allowGlobalUpdate
	return b
}

// WithQueryFields sets the queryFields option
func (b *DBBuilder) WithQueryFields(queryFields bool) *DBBuilder {
	b.config.queryFields = &queryFields
	return b
}

// WithCreateBatchSize sets the createBatchSize option
func (b *DBBuilder) WithCreateBatchSize(createBatchSize int) *DBBuilder {
	b.config.createBatchSize = &createBatchSize
	return b
}

// WithTranslateError sets the translateError option
func (b *DBBuilder) WithTranslateError(translateError bool) *DBBuilder {
	b.config.translateError = &translateError
	return b
}

// WithPropagateUnscoped sets the propagateUnscoped option
func (b *DBBuilder) WithPropagateUnscoped(propagateUnscoped bool) *DBBuilder {
	b.config.propagateUnscoped = &propagateUnscoped
	return b
}

// Open opens a new database connection
func (b *DBBuilder) Open() (*DB, error) {
	var dialect gorm.Dialector
	if b.driverName == Mysql {
		dialect = mysql.Open(b.dsn)
	} else {
		dialect = postgres.Open(b.dsn)
	}

	config := &gorm.Config{}

	db, err := gorm.Open(dialect, config)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if b.config != nil {
		pointer.SetIfNotNil(b.config.maxIdleLifeTime, func(value *time.Duration) {
			sqlDB.SetConnMaxLifetime(*value)
		})
		pointer.SetIfNotNil(b.config.maxIdleConnection, func(value *int) {
			sqlDB.SetMaxIdleConns(*value)
		})

		pointer.SetIfNotNil(b.config.maxOpenConnection, func(value *int) {
			sqlDB.SetMaxOpenConns(*value)
		})

		pointer.SetIfNotNil(b.config.skipDefaultTransaction, func(value *bool) {
			db.SkipDefaultTransaction = *value
		})

		pointer.SetIfNotNil(b.config.logger, func(value *logger.Interface) {
			db.Logger = *value
		})

		pointer.SetIfNotNil(b.config.dryRun, func(value *bool) {
			db.DryRun = *value
		})

		pointer.SetIfNotNil(b.config.prepareStmt, func(value *bool) {
			db.PrepareStmt = *value
		})

		pointer.SetIfNotNil(b.config.disableAutomaticPing, func(value *bool) {
			db.DisableAutomaticPing = *value
		})

		pointer.SetIfNotNil(b.config.disableForeignKeyConstraintWhenMigrating, func(value *bool) {
			db.DisableForeignKeyConstraintWhenMigrating = *value
		})

		pointer.SetIfNotNil(b.config.ignoreRelationshipsWhenMigrating, func(value *bool) {
			db.IgnoreRelationshipsWhenMigrating = *value
		})

		pointer.SetIfNotNil(b.config.disableNestedTransaction, func(value *bool) {
			db.DisableNestedTransaction = *value
		})

		pointer.SetIfNotNil(b.config.allowGlobalUpdate, func(value *bool) {
			db.AllowGlobalUpdate = *value
		})

		pointer.SetIfNotNil(b.config.queryFields, func(value *bool) {
			db.QueryFields = *value
		})

		pointer.SetIfNotNil(b.config.createBatchSize, func(value *int) {
			db.CreateBatchSize = *value
		})

		pointer.SetIfNotNil(b.config.translateError, func(value *bool) {
			db.TranslateError = *value
		})

		pointer.SetIfNotNil(b.config.propagateUnscoped, func(value *bool) {
			db.PropagateUnscoped = *value
		})

	}
	return &DB{db: db}, nil
}
