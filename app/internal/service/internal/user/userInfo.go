package user

import (
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/user"
)

type SInfo struct{}

var insInfo = SInfo{}

func (s *SInfo) GetUserInfo(userBasic *user.Basic, userCounter *user.Counter, id any) error {
	sqlStr := "select * from user_basic where id=?"
	err := g.MysqlDB.QueryRowx(sqlStr, id).StructScan(userBasic)
	if err != nil {
		g.Logger.Error("get user basic error", zap.Error(err))
		return err
	}
	sqlStr = "select * from user_counter where id=?"
	err = g.MysqlDB.QueryRowx(sqlStr, id).StructScan(userCounter)
	if err != nil {
		g.Logger.Error("get user counter error", zap.Error(err))
		return err
	}
	return nil
}
