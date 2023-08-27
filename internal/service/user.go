/*
 * @FilePath: /workflow-server/internal/service/user.go
 * @Author: maggot-code
 * @Date: 2023-08-14 14:57:17
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-20 19:49:15
 * @Description:
 */
package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/maggot-code/workflow-server/internal/biz"
	"github.com/maggot-code/workflow-server/pkg/errors"
	"github.com/maggot-code/workflow-server/pkg/handler"
)

var (
	NewUserError     = errors.StatusBadRequest("new user fail", "new user fail")
	ExchangeError    = errors.StatusBadRequest("wechat exchange session fail", "wechat exchange session fail")
	WechatLoginError = errors.StatusInternalServerError("wechat login fail", "wechat login fail")
)

type UserService struct {
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (us *UserService) WechatLogin(ctx *gin.Context, req *http.Request) (handler.Response, error) {
	u, err := us.uc.ExchangeSession(ctx, ctx.Param("code"))
	if err != nil {
		return nil, ExchangeError.WithReason(err.Error()).WithCause(err)
	}

	user, err := us.uc.FindUser(ctx, u)
	if err != nil {
		token, err := us.uc.LoginUser(ctx, u)
		if err != nil {
			return nil, WechatLoginError.WithReason(err.Error()).WithCause(err)
		}

		return token, nil
	}

	user.SessionKey = u.SessionKey
	user.SessionRefresh = carbon.DateTime{Carbon: carbon.Now()}
	token, err := us.uc.LoginUser(ctx, user)
	if err != nil {
		return nil, WechatLoginError.WithReason(err.Error()).WithCause(err)
	}

	return token, nil
}

func (us *UserService) GetUser(ctx *gin.Context, req *http.Request) (handler.Response, error) {
	return nil, nil
}
