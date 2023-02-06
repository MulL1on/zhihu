package view

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	g "juejin/app/global"
)

type SView struct{}

var insView SView

func (s *SView) CountView(ctx context.Context, authorId string, itemId string, itemType int) error {
	var key1, sqlStr1 string
	key2 := "user_counter"
	sqlStr2 := "select got_view_count from user_counter where user_id=?"
	field1 := fmt.Sprintf("{%s:view_count}", itemId)
	field2 := fmt.Sprintf("{%s:got_view_count}", authorId)
	switch itemType {
	case 2:
		key1 = "article_counter"
		sqlStr1 = "select view_count from article_counter where article_id=?"
	default:
		return fmt.Errorf("no such item type")
	}

	//检查是否存在article counter
	ok, err := g.Rdb.HExists(ctx, key1, field1).Result()
	if !ok {
		var viewCount int
		err = g.MysqlDB.QueryRow(sqlStr1, itemId).Scan(&viewCount)
		if err != nil {
			g.Logger.Error("get article view count from mysql error", zap.Error(err))
			return err
		}
		g.Rdb.HSet(ctx, key1, field1, viewCount)
	}

	err = g.Rdb.HIncrBy(ctx, key1, field1, 1).Err()
	if err != nil {
		g.Logger.Error("count view error", zap.Error(err))
		return err
	}

	//检查是否存在 user counter
	ok, err = g.Rdb.HExists(ctx, key2, field2).Result()
	if !ok {
		var viewCount int
		err = g.MysqlDB.QueryRow(sqlStr2, itemId).Scan(&viewCount)
		if err != nil {
			g.Logger.Error("get user view count from mysql error", zap.Error(err))
			return err
		}
		g.Rdb.HSet(ctx, key2, field2, viewCount)
	}
	err = g.Rdb.HIncrBy(ctx, key2, field2, 1).Err()
	if err != nil {
		g.Rdb.HIncrBy(ctx, key1, field1, -1)
		g.Logger.Error("count view error", zap.Error(err))
		return err
	}
	return nil
}
