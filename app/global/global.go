package global

import (
	"database/sql"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"juejin/app/internal/model/config"
)

var (
	Config  *config.Config
	Logger  *zap.Logger
	Rdb     *redis.Client
	MysqlDB *sql.DB
)
