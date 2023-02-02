package follower

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	g "juejin/app/global"
)

type SFollow struct{}

var insFollow SFollow

func (s *SFollow) DoFollow(followerId any, followeeId int64) error {
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}
	sqlStr1 := "insert into follow (follower, followee) VALUES (?,?)"
	_, err = tx.Exec(sqlStr1, followerId, followeeId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("do follower sqlStr1 error", zap.Error(err))
		return err
	}

	sqlStr2 := "update user_counter set followee_count=followee_count+1 where user_id=?"
	ret2, err := tx.Exec(sqlStr2, followerId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("do follower sqlStr2 error", zap.Error(err))
		return err
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("do follower ret2.RowAffected error", zap.Error(err))
		return err
	}

	sqlStr3 := "update user_counter set follower_count=follower_count+1 where user_id=?"
	ret3, err := tx.Exec(sqlStr3, followeeId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("do follower sqlStr3 error", zap.Error(err))
		return err
	}
	affRow3, err := ret3.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("do follower ret3.RowsAffected error", zap.Error(err))
		return err
	}

	if !(affRow3 == 1 && affRow2 == 1) {
		tx.Rollback()
		g.Logger.Error(fmt.Sprintf("affRow incorrect affRow2:%d affRow3:%d", affRow2, affRow3))
		return fmt.Errorf("internal error")
	}
	tx.Commit()
	return nil
}

func (s *SFollow) UndoFollow(followerId any, followeeId int64) error {
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}

	sqlStr1 := "delete from follow where follower=?&&followee=?"
	ret1, err := tx.Exec(sqlStr1, followerId, followeeId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("undo follow sqlStr1 error", zap.Error(err))
		return err
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("undo follow ret2.RowsAffected error", zap.Error(err))
		return err
	}
	if affRow1 != 1 {
		tx.Rollback()
		g.Logger.Error(fmt.Sprintf("undo follow affRow1 incorrect affRow1:%d", affRow1))
		return fmt.Errorf("internal error")
	}

	sqlStr2 := "update user_counter set follower_count=follower_count-1 where user_id=?"
	ret2, err := tx.Exec(sqlStr2, followeeId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("undo follow sqlStr2 error", zap.Error(err))
		return err
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("undo follow ret2.RowsAffected error", zap.Error(err))
		return err
	}

	sqlStr3 := "update user_counter set followee_count=followee_count-1 where user_id=?"
	ret3, err := g.MysqlDB.Exec(sqlStr3, followerId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("undo follow sqlStr3 error", zap.Error(err))
		return err
	}
	affRow3, err := ret3.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("undo follow ret3.RowsAffected error", zap.Error(err))
		return err
	}

	if !(affRow3 == 1 && affRow2 == 1) {
		tx.Rollback()
		g.Logger.Error(fmt.Sprintf("undo follower affRow incorrect,affRow2:%d affRow3:%d", affRow2, affRow3))
		return fmt.Errorf("internal error")
	}

	tx.Commit()
	return nil
}

func (s *SFollow) CheckIsFollowed(followeeId int64, followerId any) error {
	var id string
	sqlStr := "select id from follow where follower=?&&followee=?"
	err := g.MysqlDB.QueryRow(sqlStr, followerId, followeeId).Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			g.Logger.Error("check is follower error", zap.Error(err))
			return err
		}
		return nil
	}
	return fmt.Errorf("user is already followed")

}

func (s *SFollow) GetFollowerList(userId any, limit, pageNo int) (*[]int64, error) {
	sqlStr := "select follower from follow where followee=? order by id limit ?,?"
	rows, err := g.MysqlDB.Query(sqlStr, userId, (pageNo-1)*limit, limit)
	if err != nil {
		g.Logger.Error("get follow list error", zap.Error(err))
		return nil, err
	}

	var list = make([]int64, 0)
	defer rows.Close()
	for rows.Next() {
		var followerId int64
		err = rows.Scan(&followerId)
		if err != nil {
			return nil, err
		}
		list = append(list, followerId)
	}
	return &list, nil
}

func (s *SFollow) GetFolloweeList(userId any, limit, pageNo int) (*[]int64, error) {
	sqlStr := "select followee from follow where follower=? order by id limit ?,?"
	rows, err := g.MysqlDB.Query(sqlStr, userId, (pageNo-1)*limit, limit)
	if err != nil {
		g.Logger.Error("get follow list error", zap.Error(err))
		return nil, err
	}

	var list = make([]int64, 0)
	defer rows.Close()
	for rows.Next() {
		var followerId int64
		err = rows.Scan(&followerId)
		if err != nil {
			return nil, err
		}
		list = append(list, followerId)
	}
	return &list, nil
}
