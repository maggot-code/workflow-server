/*
 * @FilePath: /workflow-server/internal/data/user.go
 * @Author: maggot-code
 * @Date: 2023-08-14 14:52:18
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-22 17:33:47
 * @Description:
 */
package data

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maggot-code/workflow-server/internal/biz"
	"github.com/maggot-code/workflow-server/internal/conf"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	conf *conf.Bootstrap
}

func NewUserRepo(data *Data, c *conf.Bootstrap) biz.UserRepo {
	return &userRepo{data: data, conf: c}
}

func (ur *userRepo) CodeToSession(ctx *gin.Context, code string) (*biz.UserCodeToSession, error) {
	addr, err := buildCodeToSession(ur, code)
	if err != nil {
		return nil, fmt.Errorf("wechat: build code to session addr fail; %w", err)
	}

	res, err := http.Get(addr)
	if err != nil {
		return nil, fmt.Errorf("wechat: get code to session server fail; %w", err)
	}

	defer res.Body.Close()
	var ucts *biz.UserCodeToSession

	err = json.NewDecoder(res.Body).Decode(&ucts)
	if err != nil {
		return nil, fmt.Errorf("wechat: decode response to status fail; %w", err)
	}

	if ucts.Errcode != 0 {
		return nil, fmt.Errorf("wechat: code to session server error; %s", ucts.Errmsg)
	}

	return ucts, nil
}

func (ur *userRepo) FindByOpenid(ctx *gin.Context, u *biz.User) (*biz.User, error) {
	err := ur.data.db.Where("openid = ?", u.Openid).First(u).Error
	if err != nil {
		return nil, fmt.Errorf("db: find user by openid fail; %w", err)
	}

	return u, nil
}

func (ur *userRepo) Save(ctx *gin.Context, u *biz.User) (*biz.User, error) {
	if err := ur.data.db.Save(u).Error; err != nil {
		return nil, fmt.Errorf("db: save user fail; %w", err)
	}

	return u, nil
}
