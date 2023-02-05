package cache

import (
	"context"
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	g "juejin/app/global"
	"strconv"
	"strings"
	"time"
)

func ScanCounterCache(ctx context.Context, key string) {
	start := time.Now()
	iter := g.Rdb.HScan(ctx, key, 0, "", 0).Iterator()
	var data = make([]string, 0)
	for iter.Next(ctx) {
		err := iter.Err()
		if err != nil {
			g.Logger.Error(fmt.Sprintf("scan hash %s error", key), zap.Error(err))
			return
		}

		data = append(data, iter.Val())
	}
	if len(data) > 0 {
		err := updateCounter(ctx, data, key)
		if err != nil {
			g.Logger.Error("update counter error", zap.Error(err))
			return
		}
	}

	cost := time.Since(start)
	g.Logger.Info(fmt.Sprintf("scan %s cache successfully", key), zap.Duration("cost", cost))

}

func ScanCheckDiggCache(ctx context.Context) {
	start := time.Now()
	iter := g.Rdb.Scan(ctx, 0, "*:*:*", 0).Iterator()
	var data = make([]string, 0)
	for iter.Next(ctx) {
		err := iter.Err()
		if err != nil {
			g.Logger.Error(fmt.Sprintf("scan check digg error"), zap.Error(err))
			return
		}
		data = append(data, iter.Val())
		value := g.Rdb.Get(ctx, iter.Val()).Val()
		data = append(data, value)
	}
	err := updateCheckDigg(data)
	if err != nil {
		return
	}
	cost := time.Since(start)
	g.Logger.Info("scan check digg successfully", zap.Duration("cost", cost))
}

func genUserSqlStr(column string) string {
	var sqlStr string
	switch column {
	case "digg_article_count":
		sqlStr = "update user_counter set digg_article_count = ? where user_id=?"
	case "got_view_count":
		sqlStr = "update user_counter set got_view_count = ? where user_id=?"
	case "got_digg_count":
		sqlStr = "update user_counter set got_digg_count = ? where user_id=?"
	}
	return sqlStr
}

func genArticleSqlStr(column string) string {
	var sqlStr string
	switch column {
	case "digg_count":
		sqlStr = "update article_counter set digg_count=? where article_id=?"
	case "view_count":
		sqlStr = "update article_counter set view_count=? where article_id=?"
	}
	return sqlStr
}

func genCommentSqlStr(column string) string {
	var sqlStr string
	switch column {
	case "digg_count":
		sqlStr = "update comment set digg_count=?  where comment_id=?"
	}
	return sqlStr
}

func genReplySqlStr(column string) string {
	var sqlStr string
	switch column {
	case "digg_count":
		sqlStr = "update reply set digg_count=?  where reply_id=?"
	}
	return sqlStr
}

func updateCounter(ctx context.Context, data []string, key string) error {
	var gen func(string) string
	switch key {
	case "user_counter":
		gen = genUserSqlStr
	case "reply_counter":
		gen = genReplySqlStr
	case "article_counter":
		gen = genArticleSqlStr
	case "comment_counter":
		gen = genCommentSqlStr

	}

	for i := 0; i < len(data)-1; i += 2 {
		part := strings.SplitN(data[i], ":", 2)
		id := part[0][1:]
		column := part[1][:len(part[1])-1]
		count, _ := strconv.ParseInt(data[i+1], 10, 64)
		_, err := g.MysqlDB.Exec(gen(column), count, id)
		if err != nil {
			g.Logger.Error("update counter error", zap.Error(err))
			return err
		}
		err = g.Rdb.HDel(ctx, key, data[i]).Err()
		if err != nil {
			g.Logger.Warn("delete counter cache error")
		}
	}
	return nil
}

func updateCheckDigg(data []string) error {
	for i := 0; i < len(data)-1; i += 2 {
		part := strings.SplitN(data[i], ":", 3)
		userId := part[0]
		itemId := part[1]
		itemType, _ := strconv.Atoi(part[2])
		if data[i+1] == "1" {
			var id int
			err := g.MysqlDB.QueryRow("select id from digg where user_id=?&&digg.item_id=?&&digg.item_type=?", userId, itemId, itemType).Scan(&id)
			if err != nil {
				if err != sql.ErrNoRows {
					g.Logger.Error("update check digg error", zap.Error(err), zap.String("user_id", userId), zap.String("item_id", itemId), zap.Int("item_type", itemType))
					return err
				}
				_, err = g.MysqlDB.Exec("insert into digg (user_id,item_id, item_type) values (?,?,?)", userId, itemId, itemType)
				if err != nil {
					g.Logger.Error("update check digg error", zap.Error(err), zap.String("user_id", userId), zap.String("item_id", itemId), zap.Int("item_type", itemType))
					return err
				}

			}
		}
	}
	return nil
}
