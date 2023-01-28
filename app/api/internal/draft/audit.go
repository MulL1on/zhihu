package draft

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"juejin/app/internal/model/draft"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
)

type AuditApi struct{}

var insAudit AuditApi

func (a *AuditApi) CreateDraft(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	var draftSubject = &draft.Draft{}
	err := c.BindJSON(draftSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusBadRequest, "form error")
		return
	}
	err = service.Draft().Audit().CreateDraft(userId, draftSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	err = service.Draft().Audit().GetDetail(draftSubject.DraftId, draftSubject)
	if err != nil {

		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.OkWithData(c, "create draft successfully", draftSubject)
}

func (a *AuditApi) UpdateDraft(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "unauthorized")
		return
	}
	var draftSubject = &draft.Draft{}
	err := c.BindJSON(draftSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusBadRequest, "form error")
		return
	}

	err = service.Draft().Audit().CheckAuth(draftSubject.DraftId, userId)
	if err != nil {
		if err.Error() == "no such draft" {
			resp.ResponseFail(c, http.StatusOK, "no such draft")
			return
		}
		if err.Error() == "unauthorized" {
			resp.ResponseFail(c, http.StatusUnauthorized, "unauthorized")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	err = service.Draft().Audit().UpdateDraft(userId, draftSubject)

	if err != nil {
		if err.Error() == "no such draft" {
			resp.ResponseFail(c, http.StatusInternalServerError, "no such draft")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	service.Draft().Audit().GetDetail(draftSubject.DraftId, draftSubject)
	resp.OkWithData(c, "success", draftSubject)
}

func (a *AuditApi) GetDraftDetail(c *gin.Context) {
	userId, ok := c.Get("id")
	draftId := c.Query("draft_id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}

	err := service.Draft().Audit().CheckAuth(draftId, userId)
	if err != nil {
		if err.Error() == "no such draft" {
			resp.ResponseFail(c, http.StatusOK, "no such draft")
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
	resp.OkWithData(c, "get draft info successfully", draftSubject)

}

func (a *AuditApi) DeleteDraft(c *gin.Context) {
	userId, ok := c.Get("id")
	draftId := c.PostForm("draft_id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}

	err := service.Draft().Audit().CheckAuth(draftId, userId)
	if err != nil {
		if err.Error() == "no such draft" {
			resp.ResponseFail(c, http.StatusOK, "no such draft")
			return
		}
		if err.Error() == "unauthorized" {
			resp.ResponseFail(c, http.StatusUnauthorized, "unauthorized")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	err = service.Draft().Audit().DeleteDraft(draftId)
	if err != nil {

		if err.Error() == "no such draft" {
			resp.ResponseFail(c, http.StatusInternalServerError, "no such draft")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "delete draft successfully")
}
