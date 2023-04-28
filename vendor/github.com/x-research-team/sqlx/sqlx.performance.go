package sqlx

import "time"

var (
	MaxIdleTime                = time.Minute * 5
	MaxIdleConns               = 15
	MaxLifetime  time.Duration = 0
	MaxOpenConns               = 35
)
