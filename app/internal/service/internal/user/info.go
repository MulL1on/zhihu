package user

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/user"
)

type SInfo struct{}

var insInfo = SInfo{}

func (s *SInfo) GetUserInfo(ctx context.Context, userBasic *user.Basic, userCounter *user.Counter, id any) error {
	sqlStr := "select digg_article_count,digg_shortmsg_count,followee_count,follower_count,got_digg_count,got_view_count,post_article_count,post_shortmsg_count,select_online_course_count,collection_set_count from user_counter where user_id = ?"
	err := g.MysqlDB.QueryRow(sqlStr, id).Scan(
		&userCounter.DiggArticleCount,
		&userCounter.DiggShortmsgCount,
		&userCounter.FolloweeCount,
		&userCounter.FollowerCount,
		&userCounter.GotDiggCount,
		&userCounter.GotViewCount,
		&userCounter.PostArticleCount,
		&userCounter.PostShortmsgCount,
		&userCounter.SelectOnlineCourseCount,
		&userCounter.CollectionSetCount)
	if err != nil {
		g.Logger.Error("get user counter error", zap.Error(err))
		return err
	}

	err = getCounterCache(ctx, userCounter, id)
	if err != nil {
		return err
	}

	sqlStr = "select description,avatar,company,job_title from user_basic where user_id=?"
	err = g.MysqlDB.QueryRow(sqlStr, id).Scan(&userBasic.Description, &userBasic.Avatar, &userBasic.Company, &userBasic.JobTitle)
	if err != nil {
		g.Logger.Error("get user basic error", zap.Error(err))
		return err
	}
	return nil
}

func (s *SInfo) UpdateUserInfo(userBasic *user.Basic, id any) error {
	sqlStr := "update user_basic set avatar=?,description=?,company=?,job_title=? usern where user_id=?"
	_, err := g.MysqlDB.Exec(sqlStr, userBasic.Avatar, userBasic.Description, userBasic.Company, userBasic.JobTitle, id)
	if err != nil {
		g.Logger.Error("update user info error", zap.Error(err))
		return err
	}
	return nil
}

func (s *SInfo) GetUserBasic(userBasic *user.Basic, id any) error {
	sqlStr := "select description,avatar,company,job_title from user_basic where user_id=?"
	err := g.MysqlDB.QueryRow(sqlStr, id).Scan(&userBasic.Description, &userBasic.Avatar, &userBasic.Company, &userBasic.JobTitle)
	if err != nil {
		g.Logger.Error("get user basic error", zap.Error(err))
		return err
	}
	return nil
}

func getCounterCache(ctx context.Context, userCounter *user.Counter, id any) error {
	key := "user_counter"
	field1 := fmt.Sprintf("{%d:digg}", id)
	field2 := fmt.Sprintf("{%d:gotDigg}", id)
	field3 := fmt.Sprintf("{%d:gotView}", id)
	cmd := g.Rdb.HMGet(ctx, key, field1, field2, field3)
	err := cmd.Err()
	if err != nil {
		if err != redis.Nil {
			g.Logger.Error("get user counter cache error", zap.Error(err))
			return err
		}
		g.Rdb.HMSet(ctx, key, field1, userCounter.DiggArticleCount, field2, userCounter.GotDiggCount, field3, userCounter.GotViewCount)
		return nil
	}
	var cache = make([]int, 3)
	err = cmd.Scan(cache)
	if err != nil {
		g.Logger.Error("get user counter cache error", zap.Error(err))
		return err
	}
	userCounter.DiggArticleCount = cache[1]
	userCounter.GotDiggCount = cache[2]
	userCounter.GotViewCount = cache[3]
	return nil
}
