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
	"math/rand"
	"regexp"
	"time"
)

const EmailCheckRule = `^([A-Za-z0-9_/.-]+)@([0-9a-z\.-]+)\.([a-z\.]{2,6})$`
const PhoneCheckRule = `^1[3,5,8]\d{9}$`

type SUser struct{}

var insUser = SUser{}

func (s *SUser) CheckUserIsExist(username string) error {
	var name string
	sqlStr := "select username from user_auth where username=?"
	err := g.MysqlDB.QueryRow(sqlStr, username).Scan(&name)
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

func (s *SUser) CheckMailIsExist(email string) error {
	user := &user.Auth{}
	sqlStr := "select * from user_auth where email=?"
	err := g.MysqlDB.QueryRow(sqlStr, email).Scan(&user)
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

func (s *SUser) EncryptPassword(password string) (string, error) {
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encryptPassword), nil
}

func (s *SUser) CreateUser(userSubject *user.Auth) error {
	sqlStr := "insert into user_auth (username,password,email,phone,create_time) values (?,?,?,?,?)"
	_, err := g.MysqlDB.Exec(sqlStr, userSubject.Username, userSubject.Password, userSubject.Email, userSubject.Phone, time.Now())
	if err != nil {
		g.Logger.Error("create mysql record failed", zap.Error(err))
		return err
	}
	return nil
}
func (s *SUser) GetEncryptPassword(username string) (string, error) {
	var pwd string
	sqlStr := "select password from user_auth where username = ?"
	err := g.MysqlDB.QueryRow(sqlStr, username).Scan(&pwd)
	if err != nil {
		g.Logger.Error("get encrypt password failed", zap.Error(err))
		return "", err
	}
	return pwd, nil
}
func (s *SUser) CheckPassword(password, encryptPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPwd), []byte(password))
	return err == nil
}
func (s *SUser) GenerateToken(ctx context.Context, user *user.Auth) (string, error) {
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
	err = g.Rdb.Set(ctx, fmt.Sprintf("jwt:%d", user.Id), tokenString, time.Duration(config.ExpiresTime)*time.Second).Err()
	if err != nil {
		g.Logger.Error("set redis cache failed.",
			zap.Error(err), zap.String("key", "jwt:[id]"),
			zap.Int64("id", user.Id),
		)
		return "", fmt.Errorf("internal err")
	}
	return tokenString, nil

}

func (s *SUser) VerifyEmailFormat(email string) bool {
	reg := regexp.MustCompile(EmailCheckRule)
	return reg.MatchString(email)
}

func (s *SUser) VerifyPhoneFormat(phone string) bool {
	reg := regexp.MustCompile(PhoneCheckRule)
	return reg.MatchString(phone)
}

func (s *SUser) SendCode(ctx context.Context, email string) error {

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

func (s *SUser) CheckCode(ctx context.Context, email, code string) (bool, error) {
	cmd := g.Rdb.Get(ctx, fmt.Sprintf("verify_code:%s", email))
	err := cmd.Err()
	if err != nil {
		g.Logger.Error("check verify code from redis failed", zap.Error(err))
		return false, err
	}
	return code == cmd.Val(), nil
}
