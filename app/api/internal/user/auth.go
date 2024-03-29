package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/oAuth"
	"juejin/app/internal/model/user"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"juejin/utils/cookie"
	"net/http"
)

type AuthApi struct{}

var insAuth = AuthApi{}

func (a *AuthApi) Register(c *gin.Context) {
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
	err = service.User().Auth().CheckUserIsExist(userSubject.Username)
	if err != nil {
		if err.Error() == "username is already exist" {
			resp.ResponseFail(c, http.StatusBadRequest, "username is already exist")
			return
		} else {
			resp.ResponseFail(c, http.StatusInternalServerError, "check username's existence error")
			return
		}
	}
	err = service.User().Auth().CheckEmailIsExist(userSubject.Username)
	if err != nil {
		if err.Error() == "email is already exist" {
			resp.ResponseFail(c, http.StatusBadRequest, "email is already exist")
			return
		} else {
			resp.ResponseFail(c, http.StatusInternalServerError, "check email's existence error")
			return
		}
	}
	ok, err := service.User().Auth().CheckCode(c, userSubject.Email, userSubject.Code)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "check code from redis error")
		return
	}
	if !ok {
		resp.ResponseFail(c, http.StatusBadRequest, "code incorrect")
		return
	}
	userSubject.Password, err = service.User().Auth().EncryptPassword(userSubject.Password)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "encrypt password error")
		return
	}
	userSubject.Id = service.User().Auth().GenerateUid()
	err = service.User().Auth().CreateUser(userSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "create user successfully")
}

func (a *AuthApi) Login(c *gin.Context) {
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

	err = service.User().Auth().CheckUserIsExist(userSubject.Username)
	if err != nil {
		if err.Error() != "username is already exist" {
			resp.ResponseFail(c, http.StatusInternalServerError, "check username's existence error")
			return
		}
	} else {
		resp.ResponseFail(c, http.StatusBadRequest, "user doesn't exist")
		return
	}
	encryptPwd, err := service.User().Auth().GetEncryptPassword(userSubject.Username)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "get encrypt password failed")
		return
	}
	if !service.User().Auth().CheckPassword(userSubject.Password, encryptPwd) {
		resp.ResponseFail(c, http.StatusBadRequest, "invalid password or username")
		return
	}

	//获取用户id
	_ = g.MysqlDB.QueryRow("select id from user_auth where username=?", userSubject.Username).Scan(&userSubject.Id)

	//生成token
	tokenString, err := service.User().Auth().GenerateToken(userSubject)
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

func (a *AuthApi) SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if !service.User().Auth().VerifyEmailFormat(email) {
		resp.ResponseFail(c, http.StatusBadRequest, "email pattern is incorrect")
		return
	}
	err := service.User().Auth().CheckEmailIsExist(email)
	if err != nil {
		if err.Error() == "email is already exist" {
			resp.ResponseFail(c, http.StatusOK, "email is already exist")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	//判断是否频繁请求发送验证码
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

	err = service.User().Auth().SendCode(c, email)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "send code failed")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "send code successfully")
}

func (a *AuthApi) Logout(c *gin.Context) {
	var token string
	cookieConfig := g.Config.App.Cookie
	cookieWriter := cookie.NewCookieWriter(cookieConfig.Secret,
		cookie.Option{
			Config: cookieConfig.Cookie,
			Ctx:    c,
		})
	cookieWriter.Get("x-token", &token)
	err := service.User().Auth().AddTokenToBlackList(c, token)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "set redis key fail")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "log out successfully")
}

func (a *AuthApi) GithubOAuthCodeCallback(c *gin.Context) {
	code := c.Query("code")
	ac, err := service.User().Auth().GetGithubAccessToken(code)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	gUser, err := service.User().Auth().GetGithubUserinfo(ac)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	var userSubject = &user.Auth{}

	err = service.User().Auth().CheckGithubUser(gUser.GithubId)
	if err != nil {
		if err.Error() != "github id is already exist" {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
			return
		}
	} else {
		userSubject.Id = service.User().Auth().GenerateUid()
		userSubject.Username = gUser.Login
		userSubject.GithubId = gUser.GithubId
		err = service.User().Auth().CreateUser(userSubject)
		if err != nil {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
			return
		}
	}
	err = g.MysqlDB.QueryRow("select id from user_auth where github_id=?", gUser.GithubId).Scan(&userSubject.Id)
	if err != nil {
		g.Logger.Error("get user id error", zap.Error(err))
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	//生成token
	tokenString, err := service.User().Auth().GenerateToken(userSubject)
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
	resp.ResponseSuccess(c, http.StatusOK, "login successfully"+gUser.Login)
}

func (a *AuthApi) GithubGetUserInfo(c *gin.Context) {
	var githubAc = &oAuth.GithubOAuthAc{}
	err := c.BindJSON(githubAc)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
}
