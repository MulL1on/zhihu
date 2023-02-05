package digg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/digg"
	"time"
)

type SLike struct{}

var insLike SLike

func (s *SLike) DoDigg(ctx context.Context, userId any, like *digg.Like) error {
	key1 := fmt.Sprintf("%d:%s:%d", userId.(int64), like.ItemId, like.ItemType)
	var key2, sqlStr string
	key3 := "user_counter"
	field1 := fmt.Sprintf("{%s:digg_count}", like.ItemId)
	field2 := fmt.Sprintf("{%s:got_digg_count}", like.AuthorId)
	field3 := fmt.Sprintf("{%d:digg_article_count}", userId.(int64))
	err := g.Rdb.Set(ctx, key1, 1, 24*time.Hour).Err()
	if err != nil {

		g.Logger.Error("do digg error", zap.Error(err))
		return err
	}
	switch like.ItemType {
	case 2:
		key2 = "article_counter"
		sqlStr = "select digg_count from article_counter where article_id=?"
	case 5:
		key2 = "comment_counter"
		sqlStr = "select digg_count from comment where comment_id=?"
	case 6:
		key2 = "reply_counter"
		sqlStr = "select digg_count from reply where reply_id=?"

	}

	ok, _ := g.Rdb.HExists(ctx, key2, field1).Result()

	if !ok {
		var data int
		err = g.MysqlDB.QueryRow(sqlStr, like.ItemId).Scan(&data)
		if err != nil {
			g.Rdb.Del(ctx, key1)
			g.Logger.Error("do digg error", zap.Error(err))
			return err
		}
		g.Logger.Info("get data from mysql", zap.Int("data", data))

		g.Rdb.HSet(ctx, key2, field1, data)
	}

	err = g.Rdb.HIncrBy(ctx, key2, field1, 1).Err()
	if err != nil {
		g.Rdb.Del(ctx, key1)
		g.Logger.Error("do digg error", zap.Error(err))
		return err
	}

	//点赞的用户赞过+1
	ok, _ = g.Rdb.HExists(ctx, key3, field2).Result()
	if !ok {
		var data int
		err = g.MysqlDB.QueryRow("select digg_article_count from user_counter where user_id=?", userId).Scan(&data)
		if err != nil {
			g.Rdb.Del(ctx, key1)
			g.Rdb.HIncrBy(ctx, key2, field1, -1)
			g.Logger.Error("do digg error", zap.Error(err))
			return err
		}
		g.Rdb.HSet(ctx, key3, field2, data)

	}
	err = g.Rdb.HIncrBy(ctx, key3, field2, 1).Err()
	if err != nil {
		g.Rdb.Del(ctx, key1)
		g.Rdb.HIncrBy(ctx, key2, field1, -1)
		g.Logger.Error("do digg error", zap.Error(err))
		return err
	}

	//被点赞的用户获得的赞+1
	ok, _ = g.Rdb.HExists(ctx, key3, field3).Result()
	if !ok {
		var data int
		err = g.MysqlDB.QueryRow("select got_digg_count from user_counter where user_id=?", like.AuthorId).Scan(&data)
		if err != nil {
			g.Rdb.Del(ctx, key1)
			g.Rdb.HIncrBy(ctx, key2, field1, -1)
			g.Rdb.HIncrBy(ctx, key3, field2, -1)
			g.Logger.Error("do digg error", zap.Error(err))
			return err
		}
		g.Rdb.HSet(ctx, key3, field3, data)
	}
	err = g.Rdb.HIncrBy(ctx, key3, field3, 1).Err()
	if err != nil {
		g.Rdb.Del(ctx, key1)
		g.Rdb.HIncrBy(ctx, key2, field1, -1)
		g.Rdb.HIncrBy(ctx, key3, field2, -1)
		g.Logger.Error("do digg error", zap.Error(err))
		return err
	}
	return nil
}

func (s *SLike) UndoDigg(ctx context.Context, userId any, like *digg.Like) error {
	key1 := fmt.Sprintf("%d:%s:%d", userId.(int64), like.ItemId, like.ItemType)
	var key2, sqlStr string
	key3 := "user_counter"
	field1 := fmt.Sprintf("{%s:digg_count}", like.ItemId)
	field2 := fmt.Sprintf("{%s:got_digg_count}", like.AuthorId)
	field3 := fmt.Sprintf("{%d:digg_article_count}", userId.(int64))
	err := g.Rdb.Set(ctx, key1, 0, 24*time.Hour).Err()
	if err != nil {

		g.Logger.Error("undo digg error", zap.Error(err))
		return err
	}
	switch like.ItemType {
	case 2:
		key2 = "article_counter"
		sqlStr = "select digg_count from article_counter where article_id=?"
	case 5:
		key2 = "comment_counter"
		sqlStr = "select digg_count from comment where comment_id=?"
	case 6:
		key2 = "reply_counter"
		sqlStr = "select digg_count from reply where reply_id=?"

	}

	ok, _ := g.Rdb.HExists(ctx, key2, field1).Result()
	if !ok {
		var data int
		err = g.MysqlDB.QueryRow(sqlStr, like.ItemId).Scan(&data)
		if err != nil {
			g.Rdb.Del(ctx, key1)
			g.Logger.Error("undo digg error", zap.Error(err))
			return err
		}
		g.Rdb.HSet(ctx, key2, field1, data)
	}
	err = g.Rdb.HIncrBy(ctx, key2, field1, -1).Err()
	if err != nil {
		g.Rdb.Del(ctx, key1)
		g.Logger.Error("undo digg error", zap.Error(err))
		return err
	}

	//点赞的用户赞过-1
	ok, _ = g.Rdb.HExists(ctx, key3, field2).Result()
	if !ok {
		var data int
		err = g.MysqlDB.QueryRow("select digg_article_count from user_counter where user_id=?", userId).Scan(&data)
		if err != nil {
			g.Rdb.Del(ctx, key1)
			g.Rdb.HIncrBy(ctx, key2, field1, 1)
			g.Logger.Error("undo digg error", zap.Error(err))
			return err
		}
		g.Rdb.HSet(ctx, key3, field2, data)
	}
	err = g.Rdb.HIncrBy(ctx, key3, field2, -1).Err()
	if err != nil {
		g.Rdb.Del(ctx, key1)
		g.Rdb.HIncrBy(ctx, key2, field1, 1)
		g.Logger.Error("undo digg error", zap.Error(err))
		return err
	}

	//被点赞的用户获得的赞-1
	ok, _ = g.Rdb.HExists(ctx, key3, field3).Result()
	if !ok {
		var data int
		err = g.MysqlDB.QueryRow("select got_digg_count from user_counter where user_id=?", like.AuthorId).Scan(&data)
		if err != nil {
			g.Rdb.Del(ctx, key1)
			g.Rdb.HIncrBy(ctx, key2, field1, 1)
			g.Rdb.HIncrBy(ctx, key3, field2, 1)
			g.Logger.Error("undo digg error", zap.Error(err))
			return err
		}
		g.Rdb.HSet(ctx, key3, field3, data)
	}
	err = g.Rdb.HIncrBy(ctx, key3, field3, -1).Err()
	if err != nil {
		g.Rdb.Del(ctx, key1)
		g.Rdb.HIncrBy(ctx, key2, field1, 1)
		g.Rdb.HIncrBy(ctx, key3, field2, 1)
		g.Logger.Error("undo digg error", zap.Error(err))
		return err
	}
	return nil
}

func (s *SLike) CheckIsDigg(ctx context.Context, userId any, itemId string, itemType int) (bool, error) {
	key1 := fmt.Sprintf("%d:%s:%d", userId.(int64), itemId, itemType)
	var status int

	err := g.Rdb.Get(ctx, key1).Scan(&status)
	if err != nil {
		if err != redis.Nil {
			g.Logger.Error("check is digg error", zap.Error(err))
			return false, err
		}

		var id int
		err = g.MysqlDB.QueryRow("select id from digg where user_id=?&&item_id=?&&item_type=?", userId, itemId, itemType).Scan(&id)
		if err != nil {
			if err != sql.ErrNoRows {
				g.Logger.Error("check is digg from mysql error", zap.Error(err))
				return false, err
			}
			return false, nil
		}
		return true, err
	}
	return status == 1, err
}
