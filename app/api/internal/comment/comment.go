package comment

import (
	"github.com/gin-gonic/gin"
	"juejin/app/internal/model/comment"
	"juejin/app/internal/model/user"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
	"strconv"
	"time"
)

type ReviewApi struct{}

type ReplyApi struct{}

var insReview ReviewApi

var insReply ReplyApi

func (a *ReviewApi) PostComment(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	var commentSubject = &comment.Comment{}
	err := c.BindJSON(commentSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusBadRequest, "json pattern incorrect")
		return
	}
	commentSubject.CreatTime = time.Now()
	err = service.Comment().Review().PostComment(userId, commentSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	commentSubject.UserId = userId.(int64)
	resp.OkWithData(c, "post comment successfully", commentSubject)
}

func (a *ReplyApi) PostReply(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	var replySubject = &comment.ReplyBrief{}
	err := c.BindJSON(replySubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusBadRequest, "json pattern incorrect")
		return
	}
	replySubject.UserId = userId.(int64)
	replySubject.CreatTime = time.Now()
	err = service.Comment().Reply().PostReply(replySubject, userId)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.OkWithData(c, "post comment successfully", replySubject)
}

func (a *ReviewApi) DeleteComment(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	commentId := c.PostForm("comment_id")

	//检查权限
	err := service.Comment().Review().CheckAuth(userId, commentId)
	if err != nil {
		if err.Error() == "no such comment" {
			resp.ResponseFail(c, http.StatusInternalServerError, "no such comment")
			return
		} else if err.Error() == "unauthorized" {
			resp.ResponseFail(c, http.StatusUnauthorized, "unauthorized")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	err = service.Comment().Review().DeleteComment(commentId)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "delete comment successfully")
}

func (a *ReplyApi) DeleteReply(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	replyId := c.PostForm("reply_id")

	//检查权限
	err := service.Comment().Reply().CheckAuth(userId, replyId)
	if err != nil {
		if err.Error() == "no such comment" {
			resp.ResponseFail(c, http.StatusInternalServerError, "no such comment")
			return
		} else if err.Error() == "unauthorized" {
			resp.ResponseFail(c, http.StatusUnauthorized, "unauthorized")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	err = service.Comment().Reply().DeleteReply(replyId)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "delete reply successfully")
}

func (a *ReviewApi) GetCommentList(c *gin.Context) {
	//userId, ok := c.Get("id")
	//if !ok {
	//	resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
	//	return
	//}
	itemId := c.PostForm("item_id")
	itemType, _ := strconv.Atoi(c.PostForm("item_type"))
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	pageNo, _ := strconv.Atoi(c.PostForm("page_no"))

	list, err := service.Comment().Review().GetCommentIdList(itemId, itemType, limit, pageNo)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "get commit id list error")
		return
	}

	var commentList = make([]*comment.List, len(*list))
	if len(*list) == 0 {
		resp.OkWithData(c, "get comment list successfully", commentList)
		return
	}
	for k, v := range *list {
		var data = &comment.List{}
		var commentInfo = &comment.Comment{}
		var userInfo = &user.InfoPack{}
		err = service.Comment().Review().GetCommentInfo(commentInfo, v)
		if err != nil {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
			return
		}
		replyList, err := service.Comment().Reply().GetReplyInfo(commentInfo.CommentId)
		if err != nil {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
			return
		}
		err = service.User().Info().GetUserInfo(c, &userInfo.Basic, &userInfo.Counter, commentInfo.UserId)
		data.CommentInfo = *commentInfo
		data.UserInfo = *userInfo
		data.ReplyInfo = *replyList
		commentList[k] = data
	}

	resp.OkWithData(c, "get comment list successfully", commentList)

}
