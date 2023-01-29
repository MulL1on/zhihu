package user

import (
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/user"
)

type SInfo struct{}

var insInfo = SInfo{}

func (s *SInfo) GetUserInfo(userBasic *user.Basic, userCounter *user.Counter, id any) error {
	sqlStr := "select digg_article_count,digg_shortmsg_count,followee_count,follower_count,got_digg_count,got_view_count,post_article_count,post_shortmsg_count,select_online_course_count from user_counter where user_id = ?"
	err := g.MysqlDB.QueryRow(sqlStr, id).Scan(
		&userCounter.DiggArticleCount,
		&userCounter.DiggShortmsgCount,
		&userCounter.FolloweeCount,
		&userCounter.FollowerCount,
		&userCounter.GotDiggCount,
		&userCounter.GotViewCount,
		&userCounter.PostArticleCount,
		&userCounter.PostShortmsgCount,
		&userCounter.SelectOnlineCourseCount)
	if err != nil {
		g.Logger.Error("get user counter error", zap.Error(err))
		return err
	}
	sqlStr = "select description,avatar,company,job_title from user_basic where user_id=?"
	err = g.MysqlDB.QueryRow(sqlStr, id).Scan(&userBasic.Description, &userBasic.Avatar, &userBasic.Company, &userBasic.JobTitle)
	if err != nil {
		g.Logger.Error("get user basic error", zap.Error(err))
		return err
	}
	return nil
}

func (s *SInfo) UpdateUserInfo(userBasic *user.Basic, id any) error {
	sqlStr := "update user_basic set avatar=?,description=?,company=?,job_title=? where user_id=?"
	_, err := g.MysqlDB.Exec(sqlStr, userBasic.Avatar, userBasic.Description, userBasic.Company, userBasic.JobTitle, id)
	if err != nil {
		g.Logger.Error("update user info error", zap.Error(err))
		return err
	}
	return nil
}
