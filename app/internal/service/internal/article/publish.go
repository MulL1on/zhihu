package article

import (
	"fmt"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/draft"
	"strconv"
	"time"
)

type SPublish struct{}

var insPublish SPublish

func (s *SPublish) Publish(draft *draft.Draft) error {
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}
	sqlStr1 := "insert into article_major (content, brief_content,cover, title, category_id, user_id, create_time) VALUES(?,?,?,?,?,?,?)"
	ret1, err := tx.Exec(sqlStr1, draft.Content, draft.BriefContent, draft.Cover, draft.Title, draft.CategoryId, draft.UserId, time.Now())
	if err != nil {
		tx.Rollback()
		g.Logger.Error("publish article sqlStr1 error", zap.Error(err))
		return err
	}

	id, _ := ret1.LastInsertId()
	articleId := strconv.FormatInt(id, 10)
	sqlStr2 := "update item_tag set item_id=? ,item_type=2 where item_id=? &&item_tag.item_type=4"
	ret2, err := tx.Exec(sqlStr2, articleId, draft.DraftId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("insert article's tags error", zap.Error(err))
		return err
	}

	affRows2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("publish article ret2.RowsAffected error", zap.Error(err))
		return err
	}

	sqlStr3 := "delete from draft where draft_id=?"
	ret3, err := tx.Exec(sqlStr3, draft.DraftId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("insert article's tags error", zap.Error(err))
		return err
	}

	affRows3, err := ret3.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("publish article ret3.RowsAffected error", zap.Error(err))
		return err
	}

	sqlStr4 := "insert into article_counter (article_id) values (?)"
	_, err = tx.Exec(sqlStr4, articleId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("publish article sqlStr4 error", zap.Error(err))
		return err
	}

	if !(affRows2 > 0 && affRows3 == 1) {
		tx.Rollback()
		g.Logger.Error("affRow incorrect ", zap.Int64("affRows2", affRows2), zap.Int64("affRow3", affRows3))
		return fmt.Errorf("internal error")
	}
	tx.Commit()
	return nil
}
