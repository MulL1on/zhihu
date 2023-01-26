package global

import (
	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"juejin/app/internal/model/config"
)

var (
	Config  *config.Config
	Logger  *zap.Logger
	Rdb     *redis.Client
	MysqlDB *sqlx.DB
)
