package initial

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// singleton
var (
	O_SERVER_CONFIG *ServerConfig

	O_DB_PG *gorm.DB
	O_REDIS *redis.Client
)
