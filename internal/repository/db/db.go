// @Author: Imam Magribi
// @Package: db
package db

import (
	"gorm.io/gorm/logger"
	"time"
)

type NewDbConfig struct {
	// MaxIdleLifeTime sets the maximum amount of time a connection may be reused.
	maxIdleLifeTime *time.Duration
	// MaxIdleConnection sets the maximum number of connections in the idle connection pool.
	maxIdleConnection *int
	// MaxOpenConnection sets the maximum number of open connections to the database.
	maxOpenConnection *int

	// GORM perform single create, update, delete operations in transactions by default to ensure database data integrity
	// You can disable it by setting `SkipDefaultTransaction` to true
	skipDefaultTransaction *bool

	// Logger
	logger *logger.Interface

	// DryRun generate sql without execute
	dryRun *bool
	// PrepareStmt executes the given query in cached statement
	prepareStmt *bool

	// DisableAutomaticPing
	disableAutomaticPing *bool

	// DisableForeignKeyConstraintWhenMigrating
	disableForeignKeyConstraintWhenMigrating *bool

	// IgnoreRelationshipsWhenMigrating
	ignoreRelationshipsWhenMigrating *bool

	// DisableNestedTransaction disable nested transaction
	disableNestedTransaction *bool

	// AllowGlobalUpdate allow global update
	allowGlobalUpdate *bool

	// QueryFields executes the SQL query with all fields of the table
	queryFields *bool

	// CreateBatchSize default create batch size
	createBatchSize *int

	// TranslateError enabling error translation
	translateError *bool

	// PropagateUnscoped propagate Unscoped to every other nested statement
	propagateUnscoped *bool
}
