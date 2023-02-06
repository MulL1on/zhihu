package tag

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/tag"
)

type SEdit struct{}

var insEdit = SEdit{}

func (s *SEdit) Delete(itemId string, idType string) error {
	sqlStr := "delete from item_tag where item_id=?&item_tag.item_type=?"
	res, err := g.MysqlDB.Exec(sqlStr, itemId, idType)
	if rowsAffect, _ := res.RowsAffected(); rowsAffect == 0 {
		return fmt.Errorf("no record")
	}
	if err != nil {
		g.Logger.Error("delete tag error", zap.Error(err))
		return err
	}
	return nil
}

func (s *SEdit) Update(tagsIds []string, itemId, idType string) error {
	sqlStr := "insert into  item_tag  (tag_id,item_id,item_type) values (?,?,?)"
	for v, _ := range tagsIds {
		_, err := g.MysqlDB.Exec(sqlStr, v, itemId, idType)
		if err != nil {
			g.Logger.Error("insert tag error", zap.Error(err))
			return err
		}
	}

	return nil
}

func (s *SEdit) GetTagListByCategory(categoryId string) (*[]tag.Tag, error) {
	var tagList = make([]tag.Tag, 0)
	sqlStr := "select t.tag_id,t.tag_name,t.create_time from tag_info t where t.tag_id in (select tag_id from item_tag where item_id=?&&item_tag.item_type=8)"
	rows, err := g.MysqlDB.Query(sqlStr, categoryId)
	if err != nil {
		if err == sql.ErrNoRows {
			return &tagList, nil
		}
		g.Logger.Error("get category tag list error", zap.String("category_id", categoryId), zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tag = &tag.Tag{}
		err := rows.Scan(&tag.TagId, &tag.TagName, &tag.CreateTime)
		if err != nil {
			g.Logger.Error("scan tag info error when get tag list", zap.Error(err))
			return nil, err
		}
		tagList = append(tagList, *tag)
	}
	return &tagList, err
}

func (s *SEdit) GetTagListByItem(itemId string, itemType int) (*[]tag.Tag, error) {
	var tagList = make([]tag.Tag, 0)
	sqlStr := "select t.tag_id,t.tag_name,t.create_time from tag_info t where t.tag_id in (select tag_id from item_tag where item_id=?&&item_tag.item_type=?)"
	rows, err := g.MysqlDB.Query(sqlStr, itemId, itemType)
	if err != nil {
		if err == sql.ErrNoRows {
			return &tagList, nil
		}
		g.Logger.Error("get category tag list error", zap.String("item_id", itemId), zap.Int("item_type", itemType), zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tag = &tag.Tag{}
		err := rows.Scan(&tag.TagId, &tag.TagName, &tag.CreateTime)
		if err != nil {
			g.Logger.Error("scan tag info error when get tag list", zap.Error(err))
			return nil, err
		}
		tagList = append(tagList, *tag)
	}
	return &tagList, err
}

func (s *SEdit) GetTagInfo(tagId string, t *tag.Tag) error {
	sqlStr := "select tag_id, tag_name, create_time from tag_info where tag_id=?"
	err := g.MysqlDB.QueryRow(sqlStr, tagId).Scan(&t.TagId, &t.CreateTime, &t.CreateTime)
	if err != nil {
		g.Logger.Error("get tag info error", zap.Error(err))
		return err
	}
	return nil
}
