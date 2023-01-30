package article

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"juejin/app/internal/model/draft"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
)

type PublishApi struct{}

var insPublish PublishApi

func (a *PublishApi) Publish(c *gin.Context) {
	userId, ok := c.Get("id")
	draftId := c.PostForm("draft_id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}

	err := service.Draft().Audit().CheckAuth(draftId, userId)
	if err != nil {
		if err.Error() == "no such draft" {
			resp.ResponseFail(c, http.StatusInternalServerError, "no such draft")
			return
		}
		if err.Error() == "unauthorized" {
			resp.ResponseFail(c, http.StatusUnauthorized, "unauthorized")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	var draftSubject = &draft.Draft{}
	err = service.Draft().Audit().GetDetail(draftId, draftSubject)
	if err != nil {
		if err == sql.ErrNoRows {
			resp.ResponseFail(c, http.StatusInternalServerError, "no such draft")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	err = service.Article().Publish().Publish(draftSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "publish article successfully")
}
