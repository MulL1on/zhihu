package boot

import (
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	g "juejin/app/global"
)

func SnowFlakeSetup() {
	node, err := snowflake.NewNode(g.Config.Snowflake.MachineId)
	if err != nil {
		g.Logger.Fatal(" new node error", zap.Error(err))
	}
	g.SfNode = node
	g.Logger.Info("initialize snowflake successfully")
}
