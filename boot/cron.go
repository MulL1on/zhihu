package boot

import (
	"context"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/utils/common/cache"
)

func ExecuteCron(ctx context.Context) {
	spec := g.Config.Cron.ScanCounterSpec
	if spec == "" {
		spec = "@every 2h"
	}
	c := cron.New(cron.WithSeconds())

	_, err := c.AddFunc(spec, func() {
		cache.ScanCounterCache(ctx, "user_counter")

	})
	if err != nil {
		g.Logger.Fatal("start user_counter error", zap.Error(err))
	}

	_, err = c.AddFunc(spec, func() {
		cache.ScanCounterCache(ctx, "article_counter")
	})
	if err != nil {
		g.Logger.Fatal("start article_counter error", zap.Error(err))
	}

	_, err = c.AddFunc(spec, func() {
		cache.ScanCounterCache(ctx, "comment_counter")
	})
	if err != nil {
		g.Logger.Fatal("start comment_counter error", zap.Error(err))
	}

	_, err = c.AddFunc(spec, func() {
		cache.ScanCounterCache(ctx, "reply_counter")
	})
	if err != nil {
		g.Logger.Fatal("start reply_counter error", zap.Error(err))
	}

	_, err = c.AddFunc(spec, func() {
		cache.ScanCheckDiggCache(ctx)
	})
	if err != nil {
		g.Logger.Fatal("start check digg error", zap.Error(err))
	}

	c.Start()
	g.Logger.Info("initialize cron successfully", zap.String("counter spec", spec))
}
