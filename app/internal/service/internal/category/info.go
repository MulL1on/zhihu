package category

import (
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/category"
)

type SInfo struct{}

var insSInfo SInfo

func (s *SInfo) GetCategoryInfo(categoryId string, c *category.Category) error {
	sqlStr := "select category_id, category_name from category where category_id=?"
	err := g.MysqlDB.QueryRow(sqlStr, categoryId).Scan(&c.CategoryId, &c.CategoryName)
	if err != nil {
		g.Logger.Error("get category info error", zap.Error(err))
		return err
	}
	return nil
}
