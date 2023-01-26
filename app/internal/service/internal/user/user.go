package user

import (
	"context"
	"database/sql"

	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	g "juejin/app/global"
	"juejin/app/internal/model/user"
	"juejin/utils/jwt"
	myjwt "juejin/utils/jwt"
	"math/rand"
	"regexp"
	"time"
)

const EmailCheckRule = `^([A-Za-z0-9_/.-]+)@([0-9a-z\.-]+)\.([a-z\.]{2,6})$`
const PhoneCheckRule = `^1[3,5,8]\d{9}$`

type SAuth struct{}

var insAuth = SAuth{}

func (s *SAuth) CheckUserIsExist(username string) error {
	var userSubject = &user.Auth{}
	sqlStr := "select * from user_auth where username=?"
	err := g.MysqlDB.QueryRowx(sqlStr, username).StructScan(userSubject)
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
	userSubject := &user.Auth{}
	sqlStr := "select * from user_auth where email=?"
	err := g.MysqlDB.QueryRowx(sqlStr, email).StructScan(userSubject)
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
	sqlStr := "insert into user_auth (username,password,email,phone,create_time,update_time) values (?,?,?,?,?,?)"
	_, err := g.MysqlDB.Exec(sqlStr, userSubject.Username, userSubject.Password, userSubject.Email, userSubject.Phone, time.Now(), time.Now())
	if err != nil {
		g.Logger.Error("create mysql record failed", zap.Error(err))
		return err
	}
	var id int
	_ = g.MysqlDB.QueryRow("select id from user_auth where username=?", userSubject.Username).Scan(&id)
	sqlStr = "insert into user_counter (id) values (?)"
	_, err = g.MysqlDB.Exec(sqlStr, id)
	if err != nil {
		g.Logger.Error("create mysql record failed", zap.Error(err))
		return err
	}
	sqlStr = "insert into user_basic (id) values (?)"
	_, err = g.MysqlDB.Exec(sqlStr, id)
	if err != nil {
		g.Logger.Error("create mysql record failed2", zap.Error(err))
		return err
	}
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