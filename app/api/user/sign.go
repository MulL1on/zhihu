package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
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
	if userSubject.Code == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "code cannot be null")
		return
	}

	if userSubject.Email == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "email cannot be null")
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
	ok, err := service.User().User().CheckCode(c, userSubject.Email, userSubject.Code)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, userSubject.Email)
		return
	}
	if !ok {
		resp.ResponseFail(c, http.StatusBadRequest, "code incorrect")
		return
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

func (a *SignApi) SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if !service.User().User().VerifyEmailFormat(email) {
		resp.ResponseFail(c, http.StatusBadRequest, "email pattern is incorrect")
		return
	}
	err := service.User().User().CheckMailIsExist(email)
	if err != nil {
		if err.Error() == "internal err" {
			resp.ResponseFail(c, http.StatusInternalServerError, "check mail failed ")
			return
		} else if err.Error() == "email is already exist" {
			resp.ResponseFail(c, http.StatusOK, "mail is already signed")
			return
		}
	}
	err = g.Rdb.Get(c, fmt.Sprintf("verify_code:%s", email)).Err()
	if err != nil {
		if err != redis.Nil {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
			return
		}
	} else {
		resp.ResponseFail(c, http.StatusBadRequest, "send code request too much")
		return
	}
	err = service.User().User().SendCode(c, email)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "send code failed")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "send code successfully")
}

func (a *SignApi) Logout(c *gin.Context) {
	var token string
	cookieConfig := g.Config.App.Cookie
	cookieWriter := cookie.NewCookieWriter(cookieConfig.Secret,
		cookie.Option{
			Config: cookieConfig.Cookie,
			Ctx:    c,
		})
	cookieWriter.Get("x-token", &token)
	err := service.User().User().AddTokenToBlackList(c, token)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "set redis key fail")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "log out successfully")
}
