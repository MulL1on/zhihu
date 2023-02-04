package cache

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	g "juejin/app/global"
)

func ScanCache(ctx context.Context, key string) error {
	for {
		iter := g.Rdb.HScan(ctx, key, 0, "*-*_count", 0).Iterator()

		for iter.Next(ctx) {
			err := iter.Err()
			if err != nil {
				g.Logger.Error(fmt.Sprintf("scan hash %s error", key), zap.Error(err))
				return err
			}
			fmt.Println(iter.Val())
		}

	}
}
