package article

import (
	"context"
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/article"
	"juejin/app/internal/model/user"
	"strconv"
)

type SInfo struct{}

var insInfo SInfo

func (s *SInfo) GetArticle(articleId string, a *article.List) error {
	sqlStr1 := "select  brief_content,cover, title, category_id, user_id, article_id, create_time from article_major where article_id=?"
	err := g.MysqlDB.QueryRow(sqlStr1, articleId).Scan(&a.ArticleInfo.BriefContent, &a.ArticleInfo.Cover, &a.ArticleInfo.Title, &a.ArticleInfo.CategoryId, &a.ArticleInfo.UserId, &a.ArticleInfo.ArticleId, &a.ArticleInfo.PublishTime)
	if err != nil {
		if err == sql.ErrNoRows {
			g.Logger.Error(fmt.Sprintf("no such article:%v", articleId))
			return fmt.Errorf("no such article")
		}
		g.Logger.Error("get article major error", zap.Error(err))
		return err
	}
	err = getUserInfo(context.Background(), &a.AuthorInfo.Basic, &a.AuthorInfo.Counter, a.ArticleInfo.UserId)
	if err != nil {
		g.Logger.Error("get author info error when get article ", zap.Error(err))
		return err
	}

	//获取文章计数
	sqlStr2 := "select digg_count, view_count, collect_count, comment_count from article_counter where article_id =?"
	err = g.MysqlDB.QueryRow(sqlStr2, articleId).Scan(&a.ArticleInfo.DiggCount, &a.ArticleInfo.ViewCount, &a.ArticleInfo.CollectCount, &a.ArticleInfo.CommentCount)
	if err != nil {
		g.Logger.Error("get article counter from mysql error", zap.Error(err))
		return err
	}
	if err != nil {
		if err == sql.ErrNoRows {
			g.Logger.Error(fmt.Sprintf("no such article:%v", articleId))
			return fmt.Errorf("no such article")
		}
		g.Logger.Error("get article collect or comment count error", zap.Error(err))
		return err
	}
	err = getArticleCounterCache(context.Background(), &a.ArticleInfo)
	if err != nil {
		return err
	}

	//获取作者信息
	err = getUserInfo(context.Background(), &a.AuthorInfo.Basic, &a.AuthorInfo.Counter, a.ArticleInfo.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (s *SInfo) GetArticleDetail(articleId string, a *article.List) error {
	//获取文章信息
	sqlStr1 := "select content, brief_content, cover,title, category_id, user_id, article_id, create_time from article_major where article_id=?"
	err := g.MysqlDB.QueryRow(sqlStr1, articleId).Scan(&a.ArticleInfo.Content, &a.ArticleInfo.BriefContent, &a.ArticleInfo.Cover, &a.ArticleInfo.Title, &a.ArticleInfo.CategoryId, &a.ArticleInfo.UserId, &a.ArticleInfo.ArticleId, &a.ArticleInfo.PublishTime)
	if err != nil {
		if err == sql.ErrNoRows {
			g.Logger.Error(fmt.Sprintf("no such article:%v", articleId))
			return fmt.Errorf("no such article")
		}
		g.Logger.Error("get article major error", zap.Error(err))
		return err
	}

	//获取文章的计数
	sqlStr2 := "select digg_count, view_count, collect_count, comment_count from article_counter where article_id =?"
	err = g.MysqlDB.QueryRow(sqlStr2, articleId).Scan(&a.ArticleInfo.DiggCount, &a.ArticleInfo.ViewCount, &a.ArticleInfo.CollectCount, &a.ArticleInfo.CommentCount)
	if err != nil {
		if err == sql.ErrNoRows {
			g.Logger.Error(fmt.Sprintf("no such article:%v", articleId))
			return fmt.Errorf("no such article")
		}
		g.Logger.Error("get article counter from mysql error", zap.Error(err))
		return err
	}

	err = getArticleCounterCache(context.Background(), &a.ArticleInfo)
	if err != nil {
		return err
	}

	//获取作者信息
	err = getUserInfo(context.Background(), &a.AuthorInfo.Basic, &a.AuthorInfo.Counter, a.ArticleInfo.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (s *SInfo) GetArticleListByDigg(limit, pageNo int, categoryId, tagId string) (*[]string, error) {
	var articleList = make([]string, 0)
	var sqlStr string
	var rows *sql.Rows
	var err error
	if tagId != "" {
		sqlStr = "select article_id from article_counter where article_id in (select a.article_id from article_major as a where category_id=? in (select item_id from item_tag where tag_id=?&&item_tag.item_type=2)) order by digg_count desc limit ?,?"
		rows, err = g.MysqlDB.Query(sqlStr, categoryId, tagId, (pageNo-1)*limit, limit)
	} else {
		sqlStr = "select article_id from article_counter where article_id in (select a.article_id from article_major as a where category_id=? ) order by digg_count desc limit ?,?"
		rows, err = g.MysqlDB.Query(sqlStr, categoryId, (pageNo-1)*limit, limit)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return &articleList, nil
		}
		g.Logger.Error("get article list error", zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var articleId string
		err = rows.Scan(&articleId)
		if err != nil {
			g.Logger.Error("scan article id error when get article list", zap.Error(err))
			return nil, err
		}
		articleList = append(articleList, articleId)
	}
	return &articleList, nil

}

func (s *SInfo) GetArticleListByTime(limit, pageNo int, categoryId, tagId string) (*[]string, error) {
	var articleList = make([]string, 0)
	var sqlStr string
	var rows *sql.Rows
	var err error
	if tagId != "" {
		sqlStr = "select article_id from article_major where article_id in (select a.article_id from article_major as a where category_id=? in (select item_id from item_tag where tag_id=?&&item_tag.item_type=2)) order by create_time desc limit ?,?"
		rows, err = g.MysqlDB.Query(sqlStr, categoryId, tagId, (pageNo-1)*limit, limit)
	} else {
		sqlStr = "select article_id from article_major where article_id in (select a.article_id from article_major as a where category_id=? ) order by create_time desc limit ?,?"
		rows, err = g.MysqlDB.Query(sqlStr, categoryId, (pageNo-1)*limit, limit)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return &articleList, nil
		}
		g.Logger.Error("get article list error", zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var articleId string
		err = rows.Scan(&articleId)
		if err != nil {
			g.Logger.Error("scan article id error when get article list", zap.Error(err))
			return nil, err
		}
		articleList = append(articleList, articleId)
	}
	return &articleList, nil

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
		g.Logger.Error("get user counter error", zap.Error(err))
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

func getArticleCounterCache(ctx context.Context, a *article.Article) error {
	key := "article_counter"
	field1 := fmt.Sprintf("{%s:digg_count}", a.ArticleId)
	field2 := fmt.Sprintf("{%s:view_count}", a.ArticleId)
	var cache = make([]int, 2)
	ok, err := g.Rdb.HExists(ctx, key, field1).Result()
	if err != nil {
		g.Logger.Error("'check exist error", zap.Error(err))
		return err
	}
	if !ok {
		g.Rdb.HSet(ctx, key, field1, a.DiggCount)
	}
	ok, err = g.Rdb.HExists(ctx, key, field2).Result()
	if err != nil {
		g.Logger.Error("'check exist error", zap.Error(err))
		return err
	}
	if !ok {
		g.Rdb.HSet(ctx, key, field2, a.ViewCount)
	}
	res, err := g.Rdb.HMGet(ctx, key, field1, field2).Result()
	if err != nil {
		g.Logger.Error("get article counter cache error", zap.Error(err))
		return err
	}
	for k, v := range res {
		cache[k], _ = strconv.Atoi(v.(string))
	}

	a.DiggCount = cache[0]
	a.ViewCount = cache[1]
	return nil

}
