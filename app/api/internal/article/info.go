package article

import (
	"github.com/gin-gonic/gin"
	"juejin/app/internal/model/article"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
)

type InfoApi struct{}

var insInfo InfoApi

func (a *InfoApi) GetArticleDetail(c *gin.Context) {
	articleId := c.Query("article_id")

	var articleSubject = &article.Article{}
	err := service.Article().Info().GetArticleMajor(articleId, articleSubject)
	if err != nil {
		if err.Error() == "no such article" {
			resp.ResponseFail(c, http.StatusBadRequest, "no such article")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	resp.OkWithData(c, "get article detail successfully", articleSubject)

}
