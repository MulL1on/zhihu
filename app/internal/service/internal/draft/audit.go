package draft

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/draft"
	"strconv"
	"time"
)

type SAudit struct{}

var insAudit SAudit

func (s *SAudit) CreateDraft(userId any, draft *draft.Draft) error {
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}
	sqlStr1 := "insert into draft (content, brief_content, cover,title, category_id, user_id,create_time,update_time) VALUES (?,?,?,?,?,?,?,?)"
	res, err := tx.Exec(sqlStr1, draft.Content, draft.BriefContent, draft.Cover, draft.Title, draft.CategoryId, userId, time.Now(), time.Now())
	if err != nil {
		tx.Rollback()
		g.Logger.Error("create draft sqlStr1 error", zap.Error(err))
		return err
	}
	id, _ := res.LastInsertId()
	draft.DraftId = strconv.FormatInt(id, 10)
	if len(draft.TagsIds) != 0 {
		sqlStr2 := "insert into  item_tag  (tag_id,item_id,item_type) values (?,?,4)"
		for _, v := range draft.TagsIds {
			_, err = tx.Exec(sqlStr2, v, draft.DraftId)
			if err != nil {
				tx.Rollback()
				g.Logger.Error("creat draft sqlStr2 error", zap.Error(err))
				return err
			}
		}
	}
	tx.Commit()
	return nil
}

func (s *SAudit) UpdateDraft(userId any, draft *draft.Draft) error {
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}
	sqlStr1 := "update draft set content=?,brief_content=?,cover=? ,title=?,category_id=? where draft_id=?&&draft.user_id=?"
	ret1, err := tx.Exec(sqlStr1, draft.Content, draft.BriefContent, draft.Cover, draft.Title, draft.CategoryId, draft.DraftId, userId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("update draft sqlStr1 error", zap.Error(err))
		return err
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("update draft ret1.rowsAffect() error", zap.Error(err))
		return err
	}

	sqlStr2 := "delete from item_tag where item_id=?&item_type=4"
	_, err = tx.Exec(sqlStr2, draft.DraftId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("update draft sqlStr2 error", zap.Error(err))
		return err
	}

	sqlStr3 := "insert into  item_tag  (tag_id,item_id,item_type) values (?,?,4)"
	for _, v := range draft.TagsIds {
		_, err = tx.Exec(sqlStr3, v, draft.DraftId, "4")
		if err != nil {
			tx.Rollback()
			g.Logger.Error("update draft sqlStr3 error", zap.Error(err))
			return err
		}
	}
	if affRow1 != 1 {
		tx.Rollback()
		return fmt.Errorf("no such draft")
	}
	tx.Commit()
	return nil
}

func (s *SAudit) GetDetail(draftId string, draft *draft.Draft) error {
	sqlStr := "select content, brief_content, cover,title, category_id, draft_id, user_id, create_time, update_time from draft where draft_id=?"
	err := g.MysqlDB.QueryRow(sqlStr, draftId).Scan(&draft.Content, &draft.BriefContent, &draft.Cover, &draft.Title, &draft.CategoryId, &draft.DraftId, &draft.UserId, &draft.CreateTime, &draft.UpdateTime)
	if err != nil {
		g.Logger.Error("get draft detail error", zap.Error(err))
		return err
	}
	return nil
}

func (s *SAudit) DeleteDraft(draftId string) error {

	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}
	sqlStr1 := "delete from draft where draft_id=?"
	ret1, err := tx.Exec(sqlStr1, draftId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("delete draft sqlStr1 error", zap.Error(err))
		return err
	}
	affRows1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("exec delete draft ret1.rowsAffect() error", zap.Error(err))
		return err
	}
	sqlStr2 := "delete from item_tag where item_id=?&&item_tag.item_type=4"
	ret2, err := tx.Exec(sqlStr2, draftId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("delete draft sqlStr2 error", zap.Error(err))
		return err
	}
	affRows2, err := ret2.RowsAffected()
	if err != nil {
		g.Logger.Error("delete draft affRow2 error", zap.Error(err))
		return err
	}
	if affRows1 == 1 && affRows2 > 0 {
		tx.Commit()
		return nil
	} else if affRows1 == 0 || affRows2 == 0 {
		tx.Rollback()
		g.Logger.Error("affRows2 incorrect", zap.Int64("affRows2", affRows2))
		return fmt.Errorf("no such draft")
	}

	tx.Rollback()
	return fmt.Errorf("internal error")

}

func (s *SAudit) CheckAuth(draftId string, userId any) error {
	sqlStr := "select user_id from draft where draft_id=?"
	var userIdRecorded any
	err := g.MysqlDB.QueryRow(sqlStr, draftId).Scan(&userIdRecorded)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no such draft")
		}
		g.Logger.Error("check auth error", zap.Error(err))
		return err
	}
	if userId != userIdRecorded {
		return fmt.Errorf("unauthorized")
	}
	return nil
}
