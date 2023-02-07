package rank

import (
	"github.com/gin-gonic/gin"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
	"strconv"
)

type RankApi struct{}

var insRank RankApi

func (a *RankApi) GetAuthorRankings(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	pageNo, _ := strconv.Atoi(c.Query("page_no"))
	rankings, err := service.Rank().Rank().GetAuthorRankings(limit, pageNo)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.OkWithData(c, "get author ranking successfully", rankings)
}
