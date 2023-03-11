package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	g "juejin/app/global"
	"juejin/app/internal/model/oAuth"
	"juejin/app/internal/model/user"
	"juejin/utils/jwt"
	myjwt "juejin/utils/jwt"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

const EmailCheckRule = `^([A-Za-z0-9_/.-]+)@([0-9a-z\.-]+)\.([a-z\.]{2,6})$`
const PhoneCheckRule = `^1[3,5,8]\d{9}$`

type SAuth struct{}

var insAuth = SAuth{}

func (s *SAuth) CheckUserIsExist(username string) error {
	var id string
	sqlStr := "select username from user_auth where username=?"
	err := g.MysqlDB.QueryRow(sqlStr, username).Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			g.Logger.Error("query mysql record fail", zap.Error(err))
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("username is already exist")
}

func (s *SAuth) CheckEmailIsExist(email string) error {
	var id string
	sqlStr := "select id from user_auth where email=?"
	err := g.MysqlDB.QueryRow(sqlStr, email).Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			g.Logger.Error("internal error", zap.Error(err))
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("email is already exist")
}

func (s *SAuth) EncryptPassword(password string) (string, error) {
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encryptPassword), nil
}

func (s *SAuth) CreateUser(userSubject *user.Auth) error {
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		g.Logger.Error("begin trans failed", zap.Error(err))
		return err
	}
	sqlStr1 := "insert into user_auth (id,username,password,email,phone,create_time,update_time,github_id) values (?,?,?,?,?,?,?,?)"
	_, err = tx.Exec(sqlStr1, userSubject.Id, userSubject.Username, userSubject.Password, userSubject.Email, userSubject.Phone, time.Now(), time.Now(), userSubject.GithubId)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("create user sqlStr1 error", zap.Error(err))
		return err
	}

	sqlStr2 := "insert into user_counter (user_id) values (?)"
	_, err = tx.Exec(sqlStr2, userSubject.Id)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("create user sqlStr2 failed", zap.Error(err))
		return err
	}

	sqlStr3 := "insert into user_basic (user_id) values (?)"
	_, err = tx.Exec(sqlStr3, userSubject.Id)
	if err != nil {
		tx.Rollback()
		g.Logger.Error("create user sqlStr3  failed", zap.Error(err))
		return err
	}
	tx.Commit()
	return nil
}
func (s *SAuth) GetEncryptPassword(username string) (string, error) {
	var pwd string
	sqlStr := "select password from user_auth where username = ?"
	err := g.MysqlDB.QueryRow(sqlStr, username).Scan(&pwd)
	if err != nil {
		g.Logger.Error("get encrypt password failed", zap.Error(err))
		return "", err
	}
	return pwd, nil
}
func (s *SAuth) CheckPassword(password, encryptPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPwd), []byte(password))
	return err == nil
}
func (s *SAuth) GenerateToken(user *user.Auth) (string, error) {
	config := g.Config.Middleware.Jwt
	j := jwt.NewJWT(&jwt.Config{
		SecretKey:  config.SecretKey,
		ExpireTime: config.ExpiresTime,
		BufferTime: config.BufferTime,
		Issuer:     config.Issuer})
	claims := j.CreateClaims(&jwt.BaseClaims{
		Id:         user.Id,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	})
	tokenString, err := j.GenerateToken(&claims)
	if err != nil {
		g.Logger.Error("generate token failed.", zap.Error(err))
		return "", fmt.Errorf("internal err")
	}

	return tokenString, nil

}

func (s *SAuth) VerifyEmailFormat(email string) bool {
	reg := regexp.MustCompile(EmailCheckRule)
	return reg.MatchString(email)
}

func (s *SAuth) VerifyPhoneFormat(phone string) bool {
	reg := regexp.MustCompile(PhoneCheckRule)
	return reg.MatchString(phone)
}

func (s *SAuth) SendCode(ctx context.Context, email string) error {

	code := fmt.Sprintf("%05v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000))
	err := g.Rdb.Set(ctx, fmt.Sprintf("verify_code:%s", email), code, time.Second*90).Err()
	if err != nil {
		g.Logger.Error("connect to redis error", zap.Error(err))
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress("1960441553@qq.com", "MyJueJin"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "注册验证码已发送")
	m.SetBody("text/html", "您的验证码：<b>"+code+"</b>")
	d := gomail.NewDialer("smtp.qq.com", 587, "1960441553", "uddnhcmjxzmsgchf")
	err = d.DialAndSend(m)
	if err != nil {
		g.Logger.Error("connect to redis error", zap.Error(err))
		return err
	}
	return nil
}

func (s *SAuth) CheckCode(ctx context.Context, email, code string) (bool, error) {
	cmd := g.Rdb.Get(ctx, fmt.Sprintf("verify_code:%s", email))
	err := cmd.Err()
	if err != nil {
		g.Logger.Error("check verify code from redis failed", zap.Error(err))
		return false, err
	}
	return code == cmd.Val(), nil
}

func (s *SAuth) AddTokenToBlackList(ctx context.Context, token string) error {
	jwtConfig := g.Config.Middleware.Jwt
	j := myjwt.NewJWT(&myjwt.Config{SecretKey: jwtConfig.SecretKey})
	mc, err := j.ParseToken(token)
	if err != nil {
		g.Logger.Error("parse token error", zap.Error(err))
		return err
	}
	err = g.Rdb.Set(ctx, fmt.Sprintf("black_list:%s", token), "", time.Duration(mc.ExpiresAt.Unix()-time.Now().Unix())*time.Second).Err()
	if err != nil {
		g.Logger.Error("set redis key failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *SAuth) GenerateUid() int64 {
	return g.SfNode.Generate().Int64()
}

func (s *SAuth) GetGithubAccessToken(code string) (*oAuth.GithubOAuthAc, error) {
	type Data struct {
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Code         string `json:"code"`
	}

	url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=493e24c7118bf543a97c&client_secret=3c4f4d28651c6e592bfa04997559c0e16b20b263&code=%s", code)
	req, err := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	req.Header.Add("Accept", "application/json")
	response, err := client.Do(req)
	if err != nil {
		g.Logger.Error("get github OAuth access token error", zap.Error(err))
		return nil, err
	}
	var res = &oAuth.GithubOAuthAc{}
	b, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(b, res)
	if err != nil {
		g.Logger.Error("get github OAuth access token error", zap.Error(err))
		return nil, err
	}
	defer response.Body.Close()
	return res, nil
}

func (s *SAuth) GetGithubUserinfo(ac *oAuth.GithubOAuthAc) (*oAuth.GithubUserInfo, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		g.Logger.Error("get github userinfo error", zap.Error(err))
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Add("Authorization", "token "+ac.AccessToken)
	response, err := client.Do(req)
	if err != nil {
		g.Logger.Error("get github userinfo error", zap.Error(err))
		return nil, err
	}
	var gUser = &oAuth.GithubUserInfo{}
	defer response.Body.Close()
	b, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(b, gUser)
	if err != nil {
		g.Logger.Error("get github userinfo error", zap.Error(err))
		return nil, err
	}
	g.Logger.Error("github user info", zap.Any("user info", gUser))
	return gUser, nil
}

func (s *SAuth) CheckGithubUser(githubId int64) error {
	var id string
	sqlStr := "select id from user_auth where github_id=?"
	err := g.MysqlDB.QueryRow(sqlStr, githubId).Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			g.Logger.Error("internal error", zap.Error(err))
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("github id is already exist")
}
