package tag

import (
	"github.com/gin-gonic/gin"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
)

type TagApi struct{}

var insTag TagApi

func (a *TagApi) GetTagList(c *gin.Context) {
	categoryId := c.Query("category_id")
	list, err := service.Tag().Edit().GetTagListByCategory(categoryId)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.OkWithData(c, "get tag list successfully", list)
}
