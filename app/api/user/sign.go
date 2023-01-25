package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	g "juejin/app/global"
	"juejin/app/internal/model/resp"
	"juejin/app/internal/model/user"
	"juejin/app/internal/service"
	"juejin/utils/cookie"
	"net/http"
)

type SignApi struct{}

var insSign = SignApi{}

func (a *SignApi) Register(c *gin.Context) {
	var userSubject = &user.Auth{}
	err := c.BindJSON(&userSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusBadRequest, fmt.Sprintf("bind json err:%v", err))
		return
	}
	if userSubject.Username == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "username cannot be null")
		return
	}
	if userSubject.Password == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "password cannot be null")
		return
	}
	err = service.User().User().CheckUserIsExist(userSubject.Username)
	if err != nil {
		if err.Error() == "username is already exist" {
			resp.ResponseFail(c, http.StatusBadRequest, "username is already exist")
			return
		} else {
			resp.ResponseFail(c, http.StatusInternalServerError, "check username's existence error")
			return
		}
	}
	userSubject.Password, err = service.User().User().EncryptPassword(userSubject.Password)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "encrypt password error")
		return
	}
	err = service.User().User().CreateUser(userSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "create user record error")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "create user successfully")
}

func (a *SignApi) Login(c *gin.Context) {
	var userSubject = &user.Auth{}
	err := c.BindJSON(&userSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusBadRequest, fmt.Sprintf("bind json err:%v", err))
		return
	}
	if userSubject.Username == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "username cannot be null")
		return
	}
	if userSubject.Password == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "password cannot be null")
		return
	}
	err = service.User().User().CheckUserIsExist(userSubject.Username)
	if err != nil {
		if err.Error() != "username is already exist" {
			resp.ResponseFail(c, http.StatusInternalServerError, "check username's existence error")
			return
		}
	} else {
		resp.ResponseFail(c, http.StatusBadRequest, "user doesn't exist")
		return
	}
	encryptPwd, err := service.User().User().GetEncryptPassword(userSubject.Username)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "get encrypt password failed")
		return
	}
	if !service.User().User().CheckPassword(userSubject.Password, encryptPwd) {
		resp.ResponseFail(c, http.StatusBadRequest, "invalid password or username")
		return
	}
	tokenString, err := service.User().User().GenerateToken(c, userSubject)
	if err != nil {
		switch err.Error() {
		case "internal err":
			resp.ResponseFail(c, http.StatusInternalServerError, "internal err")
		}
	}
	cookieConfig := g.Config.App.Cookie

	cookieWriter := cookie.NewCookieWriter(cookieConfig.Secret,
		cookie.Option{
			Config: cookieConfig.Cookie,
			Ctx:    c,
		})
	cookieWriter.Set("x-token", tokenString)
	resp.ResponseSuccess(c, http.StatusOK, "login successfully")
}
