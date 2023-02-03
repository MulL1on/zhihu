package comment

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/comment"
	"juejin/app/internal/model/user"
	"strconv"
)

type SReview struct{}

type SReply struct{}

var insReview SReview

var insReply SReply

func (s *SReview) PostComment(userId any, c *comment.Comment) error {
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}

	sqlStr1 := "insert into comment ( user_id,comment_content, create_time, item_id,item_type) values (?,?,?,?,?)"
	ret1, err := tx.Exec(sqlStr1, userId, c.CommentContent, c.CreatTime, c.ItemId, c.ItemType)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("post comment sqlStr1 error", zap.Error(err))
		return err
	}
	insId, err := ret1.LastInsertId()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("post comment ret1.LastInsertId error", zap.Error(err))
		return err
	}
	sqlStr2 := "update article_counter set comment_count=comment_count+1 where article_id=?"
	ret2, err := tx.Exec(sqlStr2, c.ItemType)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("post comment sqlStr2 error", zap.Error(err))
		return err
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("post comment ret2.RowsAffected error", zap.Error(err))
		return err
	}
	if affRow2 != 1 {
		tx.Rollback()
		g.Logger.Error(fmt.Sprintf("post comment affRow incorrect affRow2 :%d", affRow2))
		return fmt.Errorf("inernal error")
	}
	tx.Commit()
	c.CommentId = strconv.FormatInt(insId, 10)
	return nil
}

func (s *SReview) DeleteComment(commentId string) error {
	var articleId string
	err := g.MysqlDB.QueryRow("select  item_id from comment where comment_id = ?", commentId).Scan(&articleId)
	if err != nil {
		g.Logger.Error("delete comment get article id error", zap.Error(err))
		return err
	}
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}

	sqlStr1 := "delete from comment where comment_id=?"
	ret1, err := tx.Exec(sqlStr1, commentId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("delete comment sqlStr1 error", zap.Error(err))
		return err
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("delete comment ret1.RowsAffected error", zap.Error(err))
		return err
	}

	sqlStr2 := "update article_counter set comment_count=comment_count-1 where article_id=?"
	ret2, err := tx.Exec(sqlStr2, articleId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("post comment sqlStr2 error", zap.Error(err))
		return err
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("delete comment ret2.RowsAffected error", zap.Error(err))
		return err
	}
	if !(affRow2 == 1 && affRow1 == 1) {
		tx.Rollback()
		g.Logger.Error(fmt.Sprintf("delete comment affRow incorrect affRow1:%d affRow2:%d", affRow1, affRow2))
		return fmt.Errorf("inernal error")
	}
	tx.Commit()
	return nil
}

func (s *SReply) PostReply(r *comment.ReplyBrief, userId any) error {

	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}

	sqlStr1 := "insert into reply ( user_id,reply_comment_id,parent_reply_id,reply_user_id,reply_content, create_time, item_id,item_type) values (?,?,?,?,?,?,?,?)"
	_, err = tx.Exec(sqlStr1, userId, r.ReplyToCommentId, r.ReplyToReplyId, r.ReplyToUserId, r.ReplyContent, r.CreatTime, r.ItemId, r.ItemType)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("post reply sqlStr1 error", zap.Error(err))
		return err
	}

	sqlStr2 := "update comment set reply_count=reply_count+1 where comment_id=?"
	ret2, err := tx.Exec(sqlStr2, r.ReplyToCommentId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("post reply sqlStr2 error", zap.Error(err))
		return err
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("post reply ret2.RowsAffected error", zap.Error(err))
		return err
	}
	if affRow2 != 1 {
		tx.Rollback()
		g.Logger.Error(fmt.Sprintf("post reply affRow incorrect affRow2 :%d", affRow2))
		return fmt.Errorf("inernal error")
	}
	tx.Commit()
	return nil
}

func (s *SReply) DeleteReply(replyId string) error {
	var commentId string
	err := g.MysqlDB.QueryRow("select  reply_comment_id from reply where reply_id=?", replyId).Scan(&commentId)
	if err != nil {
		g.Logger.Error("delete reply get comment id error", zap.Error(err))
		return err
	}
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}

	sqlStr1 := "delete from reply where reply_id=?"
	ret1, err := tx.Exec(sqlStr1, replyId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("delete reply sqlStr1 error", zap.Error(err))
		return err
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("delete reply ret1.RowsAffected error", zap.Error(err))
		return err
	}

	sqlStr2 := "update comment set reply_count=reply_count-1 where comment_id=?"
	ret2, err := tx.Exec(sqlStr2, commentId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("delete reply sqlStr2 error", zap.Error(err))
		return err
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		g.Logger.Error("delete reply ret2.RowsAffected error", zap.Error(err))
		return err
	}
	if !(affRow2 == 1 && affRow1 == 1) {
		tx.Rollback()
		g.Logger.Error(fmt.Sprintf("post comment affRow incorrect affRow1:%d affRow2:%d", affRow1, affRow2))
		return fmt.Errorf("inernal error")
	}
	tx.Commit()
	return nil
}

func (s *SReview) GetCommentInfo(c *comment.Comment, commentId string) error {
	sqlStr := "select comment_id, comment_content, digg_count, reply_count, create_time, item_id, item_type, user_id from comment where comment_id=?"
	err := g.MysqlDB.QueryRow(sqlStr, commentId).Scan(&c.CommentId, &c.CommentContent, &c.DiggCount, &c.ReplyCount, &c.CreatTime, &c.ItemId, &c.ItemType, &c.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no such comment")
		}
		g.Logger.Error("get comment info error", zap.Error(err))
		return err
	}
	return nil
}

func (s *SReply) GetReplyInfo(commentId string) (*[]comment.ReplyInfo, error) {
	sqlStr1 := "select reply_id, reply_comment_id, reply_content, user_id, item_id, item_type, digg_count, create_time, parent_reply_id, reply_user_id from reply where reply_comment_id=?"
	rows, err := g.MysqlDB.Query(sqlStr1, commentId)
	if err != nil {
		g.Logger.Error("get reply info error", zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	var list = make([]comment.ReplyInfo, 0)
	for rows.Next() {
		var r comment.ReplyInfo
		err = rows.Scan(&r.ReplyInfo.ReplyId, &r.ReplyInfo.ReplyToCommentId, &r.ReplyInfo.ReplyContent, &r.ReplyInfo.UserId, &r.ReplyInfo.ItemId, &r.ReplyInfo.ItemType, &r.ReplyInfo.DiggCount, &r.ReplyInfo.CreatTime, &r.ReplyInfo.ReplyToReplyId, &r.ReplyInfo.ReplyToUserId)
		if err != nil {
			if err == sql.ErrNoRows {
				return &list, nil
			}
			g.Logger.Error("get reply list error", zap.Error(err))
			return nil, err
		}
		if r.ReplyInfo.ReplyToReplyId != "0" {
			err = g.MysqlDB.QueryRow(sqlStr1).Scan(&r.ParentReplyInfo.ReplyId, &r.ParentReplyInfo, &r.ParentReplyInfo.ReplyContent, &r.ParentReplyInfo.UserId, &r.ParentReplyInfo.ItemId, &r.ParentReplyInfo.ItemType, &r.ParentReplyInfo.DiggCount, &r.ParentReplyInfo.CreatTime, &r.ParentReplyInfo.ReplyToReplyId, &r.ParentReplyInfo.ReplyToUserId)
			if err != nil {
				g.Logger.Error("get parent reply error ", zap.Error(err))
				return nil, err
			}
		}

		err = GetUserInfo(&r.UserInfo.Basic, &r.UserInfo.Counter, r.ReplyInfo.UserId)
		if err != nil {
			return nil, err
		}
		if r.ReplyInfo.ReplyToUserId != "0" {
			err = GetUserInfo(&r.ReplyToUserInfo.Basic, &r.ReplyToUserInfo.Counter, r.ReplyInfo.ReplyToUserId)
			if err != nil {
				return nil, err
			}
		}

		list = append(list, r)
	}

	return &list, nil
}

func (s *SReply) CheckAuth(userId any, replyId string) error {
	sqlStr := "select user_id from reply where reply_id=?"
	var userIdRecorded any
	err := g.MysqlDB.QueryRow(sqlStr, replyId).Scan(&userIdRecorded)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no such reply")
		}
		g.Logger.Error("check edit comment auth error", zap.Error(err))
		return err
	}
	if userId != userIdRecorded {
		return fmt.Errorf("unauthorized")
	}
	return nil
}

func (s *SReview) CheckAuth(userId any, commentId string) error {
	sqlStr := "select user_id from comment where comment_id=?"
	var userIdRecorded any
	err := g.MysqlDB.QueryRow(sqlStr, commentId).Scan(&userIdRecorded)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no such comment")
		}
		g.Logger.Error("check edit comment auth error", zap.Error(err))
		return err
	}
	if userId != userIdRecorded {
		return fmt.Errorf("unauthorized")
	}
	return nil
}

func (s *SReview) GetCommentIdList(itemId string, itemType, limit, pageNo int) (*[]string, error) {
	sqlStr := "select comment_id from comment where item_id=?&&comment.item_type=? order by create_time limit ?,?"
	rows, err := g.MysqlDB.Query(sqlStr, itemId, itemType, (pageNo-1)*limit, limit)
	if err != nil {
		g.Logger.Error("get comment id list error", zap.Error(err))
		return nil, err
	}

	var list = make([]string, 0)
	defer rows.Close()
	for rows.Next() {
		var commentId string
		err = rows.Scan(&commentId)
		if err != nil {
			return nil, err
		}
		list = append(list, commentId)
	}
	return &list, nil
}

func GetUserInfo(userBasic *user.Basic, userCounter *user.Counter, id any) error {
	sqlStr := "select digg_article_count,digg_shortmsg_count,followee_count,follower_count,got_digg_count,got_view_count,post_article_count,post_shortmsg_count,select_online_course_count,collection_set_count from user_counter where user_id = ?"
	err := g.MysqlDB.QueryRow(sqlStr, id).Scan(
		&userCounter.DiggArticleCount,
		&userCounter.DiggShortmsgCount,
		&userCounter.FolloweeCount,
		&userCounter.FollowerCount,
		&userCounter.GotDiggCount,
		&userCounter.GotViewCount,
		&userCounter.PostArticleCount,
		&userCounter.PostShortmsgCount,
		&userCounter.SelectOnlineCourseCount,
		&userCounter.CollectionSetCount)
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
