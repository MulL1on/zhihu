package user

import (
	"github.com/gin-gonic/gin"
	"juejin/app/internal/model/resp"
	"juejin/app/internal/model/user"
	"juejin/app/internal/service"
	"net/http"
)

type InfoApi struct{}

var insInfo = InfoApi{}

func (a *InfoApi) GetUserInfo(c *gin.Context) {
	id, _ := c.Get("id")

	var userBasic = &user.Basic{}
	var userCounter = &user.Counter{}
	err := service.User().Info().GetUserInfo(userBasic, userCounter, id)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	var infoPack = user.InfoPack{
		Counter: *userCounter,
		Basic:   *userBasic,
	}
	resp.OkWithData(c, "get user info successfully", infoPack)
}
