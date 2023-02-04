package view

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	g "juejin/app/global"
)

type SView struct{}

var insView SView

func (s *SView) CountView(ctx context.Context, authorId string, itemId string, itemType int) error {
	var key1 string
	key2 := "user_counter"
	field1 := fmt.Sprintf("{%s:view}", itemId)
	field2 := fmt.Sprintf("{%s:gotView}", authorId)
	switch itemType {
	case 2:
		key1 = "article_counter"
	default:
		return fmt.Errorf("no such item type")
	}
	err := g.Rdb.HIncrBy(ctx, key1, field1, 1).Err()
	if err != nil {
		g.Logger.Error("count view error", zap.Error(err))
		return err
	}
	err = g.Rdb.HIncrBy(ctx, key2, field2, 1).Err()
	if err != nil {
		g.Rdb.HIncrBy(ctx, key1, field1, -1)
		g.Logger.Error("count view error", zap.Error(err))
		return err
	}
	return nil
}
