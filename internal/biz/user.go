/*
 * @FilePath: /workflow-server/internal/biz/user.go
 * @Author: maggot-code
 * @Date: 2023-08-14 14:52:39
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-20 20:25:17
 * @Description:
 */
package biz

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/maggot-code/workflow-server/internal/conf"
	"github.com/maggot-code/workflow-server/pkg/jwt"
	"github.com/maggot-code/workflow-server/pkg/orm"
)

type WechatError struct {
	Errmsg  string `json:"errmsg"`
	Errcode int32  `json:"errcode"`
}

type User struct {
	orm.Model
	Unionid        string          `gorm:"column:unionid" json:"unionid,omitempty"`
	Openid         string          `gorm:"column:openid" json:"openid"`
	SessionKey     string          `gorm:"column:session_key" json:"session_key"`
	SessionRefresh carbon.DateTime `gorm:"column:session_refresh" json:"session_refresh"`
}

type UserCodeToSession struct {
	User
	WechatError
}

type UserRepo interface {
	CodeToSession(ctx *gin.Context, code string) (*UserCodeToSession, error)
	FindByOpenid(ctx *gin.Context, u *User) (*User, error)
	Save(ctx *gin.Context, u *User) (*User, error)
}

type UserUseCase struct {
	conf *conf.Bootstrap
	repo UserRepo
}

func NewUserUseCase(c *conf.Bootstrap, ur UserRepo) *UserUseCase {
	return &UserUseCase{conf: c, repo: ur}
}

func (uc *UserUseCase) NewUser(ctx *gin.Context) (*User, error) {
	var user *User

	metadata, exists := ctx.Get("metadata")
	if !exists {
		return nil, fmt.Errorf("biz: get metadata error")
	}

	jsonStr, err := json.Marshal(metadata)
	if err != nil {
		return nil, fmt.Errorf("biz: json marshal error; %w", err)
	}

	if err := json.Unmarshal(jsonStr, &user); err != nil {
		return nil, fmt.Errorf("biz: json unmarshal error; %w", err)
	}

	return user, nil
}

func (uc *UserUseCase) ExchangeSession(ctx *gin.Context, code string) (*User, error) {
	ucts, err := uc.repo.CodeToSession(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("biz: exchange session error; %w", err)
	}

	return &User{
		Unionid:        ucts.Unionid,
		Openid:         ucts.Openid,
		SessionKey:     ucts.SessionKey,
		SessionRefresh: carbon.DateTime{Carbon: carbon.Now()},
	}, nil
}

func (uc *UserUseCase) FindUser(ctx *gin.Context, u *User) (*User, error) {
	user, err := uc.repo.FindByOpenid(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("biz: find user error; %w", err)
	}

	return user, nil
}

func (uc *UserUseCase) LoginUser(ctx *gin.Context, u *User) (string, error) {
	user, err := uc.repo.Save(ctx, u)
	if err != nil {
		return "", fmt.Errorf("biz: save user error; %w", err)
	}

	token := jwt.New(ctx.Request.Host, time.Now().Add(2*time.Hour), user)
	return jwt.Issue(uc.conf.Wechat.Secret, token)
}
