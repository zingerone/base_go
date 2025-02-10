// @Author: Imam Magribi
// @Package: db
package db

import (
	"time"
)

type DbConfig struct {
	maxIdleLifeTime   *time.Duration
	maxIdleConnection *int
	maxOpenConnection *int
}
