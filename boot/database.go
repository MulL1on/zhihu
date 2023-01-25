package boot

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	g "juejin/app/global"
	"time"
)

func MysqlDBSetup() {
	config := g.Config.Database.Mysql
	db, err := sql.Open("mysql", config.GetDsn())
	if err != nil {
		g.Logger.Fatal("initialize mysql failed.", zap.Error(err))
	}

	db.SetConnMaxIdleTime(g.Config.Database.Mysql.GetConnMaxIdleTime())
	db.SetConnMaxLifetime(g.Config.Database.Mysql.GetConnMaxLifeTime())
	db.SetMaxIdleConns(g.Config.Database.Mysql.MaxIdleConns)
	db.SetMaxOpenConns(g.Config.Database.Mysql.MaxOpenConns)
	err = db.Ping()
	if err != nil {
		g.Logger.Fatal("initialize mysql failed", zap.Error(err))
	}
	g.MysqlDB = db
	g.Logger.Info("initialize mysql successfully")
}

func RedisSetup() {
	config := g.Config.Database.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       config.Db,
		PoolSize: config.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		g.Logger.Fatal("connect to redis failed", zap.Error(err))
	}
	g.Rdb = rdb
	g.Logger.Info("initialize redis successfully.")
}
