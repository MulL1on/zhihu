package article

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/article"
)

type SInfo struct{}

var insInfo SInfo

func (s *SInfo) GetArticleMajor(articleId string, a *article.Article) error {
	sqlStr := "select content, brief_content, title, category_id, user_id, article_id, create_time from article_major where article_id=?"
	err := g.MysqlDB.QueryRow(sqlStr, articleId).Scan(&a.Content, &a.BriefContent, &a.Title, &a.CategoryId, &a.UserId, &a.ArticleId, &a.PublishTime)
	if err != nil {
		if err == sql.ErrNoRows {
			g.Logger.Error(fmt.Sprintf("no such article:%v", articleId))
			return fmt.Errorf("no such article")
		}
		g.Logger.Error("get article major error", zap.Error(err))
		return err
	}
	return nil
}

func (s *SInfo) GetList(articleId string, a *article.Brief) error {
	sqlStr := "select  brief_content, title, category_id, user_id, article_id, create_time from article_major"
	err := g.MysqlDB.QueryRow(sqlStr, articleId).Scan(&a.ArticleInfo.BriefContent,
		&a.ArticleInfo.Title,
		&a.ArticleInfo.CategoryId,
		&a.ArticleInfo.UserId,
		&a.ArticleInfo.ArticleId,
		&a.ArticleInfo.PublishTime)
	if err != nil {
		g.Logger.Error("get article list error", zap.Error(err))
		return err
	}
	return nil
}
