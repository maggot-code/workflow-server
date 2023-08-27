/*
 * @FilePath: /workflow-server/internal/router/v1/user.go
 * @Author: maggot-code
 * @Date: 2023-08-14 15:14:55
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-17 10:31:56
 * @Description:
 */
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/maggot-code/workflow-server/internal/conf"
	"github.com/maggot-code/workflow-server/internal/middleware"
	"github.com/maggot-code/workflow-server/internal/service"
	"github.com/maggot-code/workflow-server/pkg/handler"
)

type UserGateway struct {
	conf *conf.Bootstrap
	us   *service.UserService
}

func NewUserGateway(c *conf.Bootstrap, us *service.UserService) *UserGateway {
	return &UserGateway{conf: c, us: us}
}

func (ug *UserGateway) Register(r *gin.Engine) {
	lr := r.Group("/login")
	lr.GET("wx/:code", handler.JSON(ug.us.WechatLogin))

	ur := r.Group("/v1/user", middleware.Authentication(ug.conf))
	ur.GET("", handler.JSON(ug.us.GetUser))
}
