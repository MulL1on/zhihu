package collection

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/collection"
	"time"
)

type SEdit struct{}

var insEdit SEdit

func (s *SEdit) CreateCollectionSet(userId any, cs *collection.Set) error {
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}
	sqlStr1 := "insert into collection_set (collection_id, user_id, collection_name, description, permisssion, create_time,update_time) VALUES (?,?,?,?,?,?,?)"
	_, err = tx.Exec(sqlStr1, cs.CollectionId, userId, cs.CollectionName, cs.Description, cs.Permission, time.Now(), time.Now())
	if err != nil {
		tx.Rollback()
		g.Logger.Error("create collection set sqlStr1 err", zap.Error(err))
		return err
	}
	sqlStr2 := "update user_counter set collection_set_count=collection_set_count+1 where user_id=?"
	ret2, err := g.MysqlDB.Exec(sqlStr2, userId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("creat collection set sqlStr2 err", zap.Error(err))
		return err
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("creat collection set ret2.RowsAffected err", zap.Error(err))
		return err
	}
	if affRow2 != 1 {
		tx.Rollback()
		return fmt.Errorf("internal error")
	}
	tx.Commit()
	return nil
}

func (s *SEdit) AddSelectArticle(collectionId int64, articleId string) error {
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}
	sqlStr1 := "insert into collection_select_articles (article_id, collection_id) values (?,?)"
	_, err = tx.Exec(sqlStr1, articleId, collectionId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("add select article sqlStr1 error", zap.Error(err))
		return err
	}
	sqlStr2 := "update article_counter set collec_count=collec_count+1 where article_id=? "
	ret2, err := tx.Exec(sqlStr2, articleId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("add select article sqlStr2 error", zap.Error(err))
		return err
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("add select article affRow2.RowsAffected error", zap.Error(err))
		return err
	}
	sqlStr3 := "update collection_set set post_article_count=post_article_count+1 where collection_id=? "
	ret3, err := tx.Exec(sqlStr3, collectionId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("add select article sqlStr3 error", zap.Error(err))
		return err
	}
	affRow3, err := ret3.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("add select article affRow3.RowsAffected error", zap.Error(err))
		return err
	}
	if !(affRow2 == 1 && affRow3 == 1) {
		tx.Rollback()
		g.Logger.Error("rows affected incorrect")
		return fmt.Errorf("internal error")
	}
	tx.Commit()
	return nil

}

// CheckViewAuth 获取查看收藏夹权限 1-public 0-private
func (s *SEdit) CheckViewAuth(collectionId int64, userId any) error {
	sqlStr := "select user_id,permisssion from collection_set where collection_id=?"
	var userIdRecorded any
	var permissionRecorded int
	err := g.MysqlDB.QueryRow(sqlStr, collectionId).Scan(&userIdRecorded, &permissionRecorded)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no such collectionSet")
		}
		g.Logger.Error("check auth error", zap.Error(err))
		return err
	}
	if permissionRecorded == 0 {
		if userId != userIdRecorded {
			return fmt.Errorf("unauthorized")
		}
		return nil
	}

	return nil
}

func (s *SEdit) CheckEditAuth(collectionId int64, userId any) error {

	sqlStr := "select user_id from collection_set where collection_id=?"
	var userIdRecorded any
	err := g.MysqlDB.QueryRow(sqlStr, collectionId).Scan(&userIdRecorded)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no such collectionSet")
		}
		g.Logger.Error("check auth error", zap.Error(err))
		return err
	}
	if userId != userIdRecorded {
		return fmt.Errorf("unauthorized")
	}
	return nil
}

func (s *SEdit) GenerateUid() int64 {
	return g.SfNode.Generate().Int64()
}

func (s *SEdit) GetSelectArticleId(collectionId int64, limit, pageNo int) (*[]string, error) {
	sqlStr1 := "select article_id from collection_select_articles where collection_id=? order by id limit ?,?"
	rows, err := g.MysqlDB.Query(sqlStr1, collectionId, (pageNo-1)*limit, limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		g.Logger.Error("get collection detail: get article id list error", zap.Error(err))
		return nil, err
	}
	var list = make([]string, 0)
	for rows.Next() {
		var articleId string
		err = rows.Scan(&articleId)
		if err != nil {
			g.Logger.Error("get collection detail error", zap.Error(err))
			return nil, err
		}
		list = append(list, articleId)
	}
	return &list, nil
}

func (s *SEdit) GetCollectionInfo(collectionId int64, cs *collection.Set) error {
	sqlStr := "select collection_id, user_id, collection_name, description, permisssion, create_time, update_time, post_article_count   from collection_set where collection_id=?"
	err := g.MysqlDB.QueryRow(sqlStr, collectionId).Scan(&cs.CollectionId, &cs.UserId, &cs.CollectionName, &cs.Description, &cs.Permission, &cs.CreateTime, &cs.UpdateTime, &cs.PostArticleCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no such collection set")
		}
		g.Logger.Error("get collection info error", zap.Error(err))
		return err
	}
	return nil
}

func (s *SEdit) CheckArticleIsExist(collectionId int64, articleId string) error {
	var id string
	sqlStr := "select id from collection_select_articles where collection_id=?&&article_id=?"
	err := g.MysqlDB.QueryRow(sqlStr, collectionId, articleId).Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			g.Logger.Error("internal error", zap.Error(err))
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("article is already exist")
}

func (s *SEdit) ModifyCollectionSet(collectionId int64, cs *collection.Set) error {
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}
	sqlStr := "update collection_set set collection_name=?,description=?,permisssion=?where collection_id=?"
	ret, err := tx.Exec(sqlStr, cs.CollectionName, cs.Description, cs.Permission)
	affRow, err := ret.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("modify collection set error", zap.Error(err))
		return err
	}
	if affRow != 1 {
		tx.Rollback()
		g.Logger.Error(fmt.Sprintf("modify collection set affRow incorrect,affrow:%d,colleciton id:%d", affRow, collectionId))
		return fmt.Errorf("internal error")
	}
	tx.Commit()
	return nil
}

func (s *SEdit) DeleteCollectionSet(collectionId int64, userId any) error {
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}
	sqlStr1 := "delete  from collection_set where collection_id=?"
	ret1, err := tx.Exec(sqlStr1, collectionId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("delete collection set sqlStr2 error", zap.Error(err))
		return err
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("delete collection set ret1.RowsAffected error", zap.Error(err))
		return err
	}

	sqlStr2 := "update user_counter set collection_set_count=collection_set_count-1 where user_id=?"
	ret2, err := tx.Exec(sqlStr2, userId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("delete collection set sqlStr2 error", zap.Error(err))
		return err
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("delete collection set ret1.RowsAffected error", zap.Error(err))
		return err
	}

	if !(affRow1 == 1 && affRow2 == 1) {
		tx.Rollback()
		g.Logger.Error(fmt.Sprintf("delete collection set affRow incorrect,affrow1:%d,affrow2:%d,colleciton id:%d", affRow1, affRow2, collectionId))
		return fmt.Errorf("internal error")
	}
	tx.Commit()
	return nil
}

func (s *SEdit) RemoveSelectedArticle(articleId string, collectionId int64) error {
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}

	sqlStr1 := "delete from collection_select_articles where article_id=?&&collection_id=?"
	ret1, err := tx.Exec(sqlStr1, articleId, collectionId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("remove selected article sqlStr2 error", zap.Error(err))
		return err
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("remove selected article ret1.RowsAffected error", zap.Error(err))
		return err
	}

	sqlStr2 := "update collection_set set post_article_count=post_article_count-1 where collection_id=?"
	ret2, err := tx.Exec(sqlStr2, collectionId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("remove selected article sqlStr2 error", zap.Error(err))
		return err
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("remove selected article ret1.RowsAffected error", zap.Error(err))
		return err
	}

	if !(affRow1 == 1 && affRow2 == 1) {
		tx.Rollback()
		g.Logger.Error(fmt.Sprintf("remove selected article affRow incorrect,affrow1:%d,affrow2:%d,colleciton id:%d", affRow1, affRow2, collectionId))
		return fmt.Errorf("internal error")
	}
	tx.Commit()
	return nil

}
