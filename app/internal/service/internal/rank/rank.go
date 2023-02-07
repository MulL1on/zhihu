package rank

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/user"
	"strconv"
)

type SRank struct{}

var insSRank SRank

func (s *SRank) GetAuthorRankings(limit int, pageNo int) (*[]user.InfoPack, error) {
	sqlStr1 := "select user_id,got_view_count ,got_digg_count from user_counter limit ?,?"
	var list = make([]user.InfoPack, 0)
	var rankings = make(map[string]float32)
	rows, err := g.MysqlDB.Query(sqlStr1, (pageNo-1)*limit, limit)
	if err != nil {
		g.Logger.Error("get user ranking count error")
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var userId string
		var gotDiggCount, gotViewCount float32
		err = rows.Scan(&userId, &gotDiggCount, &gotViewCount)
		if err != nil {
			g.Logger.Error("get author rankings scan error", zap.Error(err))
			return nil, err
		}
		score := gotDiggCount*0.05 + gotDiggCount
		rankings[userId] = score
	}
	g.Logger.Error("rankings", zap.Any("rankings", rankings))
	for k, _ := range rankings {
		var u user.InfoPack
		err = getUserInfo(context.Background(), &u.Basic, &u.Counter, k)
		if err != nil {
			return nil, err
		}
		list = append(list, u)
	}

	return &list, nil
}

func getUserInfo(ctx context.Context, userBasic *user.Basic, userCounter *user.Counter, id any) error {
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
		g.Logger.Error("get user counter error", zap.Error(err), zap.Any("user_id", id))
		return err
	}

	err = getUserCounterCache(ctx, userCounter, id)
	if err != nil {
		return err
	}

	sqlStr = "select username,description,avatar,company,job_title from user_basic where user_id=?"
	err = g.MysqlDB.QueryRow(sqlStr, id).Scan(&userBasic.Username, &userBasic.Description, &userBasic.Avatar, &userBasic.Company, &userBasic.JobTitle)
	if err != nil {
		g.Logger.Error("get user basic error", zap.Error(err))
		return err
	}
	return nil
}

func getUserCounterCache(ctx context.Context, u *user.Counter, id any) error {
	key := "user_counter"
	field1 := fmt.Sprintf("{%s:digg_article_count}", id)
	field2 := fmt.Sprintf("{%s:got_digg_count}", id)
	field3 := fmt.Sprintf("{%s:got_view_count}", id)
	ok, err := g.Rdb.HExists(ctx, key, field1).Result()
	if err != nil {
		g.Logger.Error("'check exist error", zap.Error(err))
		return err
	}
	if !ok {
		g.Rdb.HSet(ctx, key, field1, u.DiggArticleCount)
	}

	ok, err = g.Rdb.HExists(ctx, key, field2).Result()
	if err != nil {
		g.Logger.Error("'check exist error", zap.Error(err))
		return err
	}
	if !ok {
		g.Rdb.HSet(ctx, key, field2, u.GotDiggCount)
	}

	ok, err = g.Rdb.HExists(ctx, key, field3).Result()
	if err != nil {
		g.Logger.Error("'check exist error", zap.Error(err))
		return err
	}
	if !ok {
		g.Rdb.HSet(ctx, key, field3, u.GotDiggCount)
	}

	res, err := g.Rdb.HMGet(ctx, key, field1, field2, field3).Result()
	if err != nil {
		g.Logger.Error("get user counter cache error", zap.Error(err))
		return err
	}
	var cache = make([]int, 3)
	for k, v := range res {
		cache[k], _ = strconv.Atoi(v.(string))
	}
	u.DiggArticleCount = cache[0]
	u.GotDiggCount = cache[1]
	u.GotViewCount = cache[2]
	return nil
}
