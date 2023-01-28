package tag

import (
	"fmt"
	"go.uber.org/zap"
	g "juejin/app/global"
)

type SEdit struct{}

var insEdit = SEdit{}

func (s *SEdit) Delete(contextId string, idType string) error {
	sqlStr := "delete from tag where context_id=?&id_type=?"
	res, err := g.MysqlDB.Exec(sqlStr, contextId, idType)
	if rowsAffect, _ := res.RowsAffected(); rowsAffect == 0 {
		return fmt.Errorf("no record")
	}
	if err != nil {
		g.Logger.Error("delete tag error", zap.Error(err))
		return err
	}
	return nil
}

func (s *SEdit) Update(tagsIds []string, contextId, idType string) error {
	sqlStr := "insert into  tag  (tag_id,context_id,id_type) values (?,?,?)"
	for v, _ := range tagsIds {
		_, err := g.MysqlDB.Exec(sqlStr, v, contextId, idType)
		if err != nil {
			g.Logger.Error("insert tag error", zap.Error(err))
			return err
		}
	}

	return nil
}
