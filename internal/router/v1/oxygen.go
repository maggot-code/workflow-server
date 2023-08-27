/*
 * @FilePath: /workflow-server/internal/router/v1/oxygen.go
 * @Author: maggot-code
 * @Date: 2023-08-16 18:40:28
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-27 21:47:24
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

type OxygenGateway struct {
	conf *conf.Bootstrap
	os   *service.OxygenService
}

func NewOxygenGateway(c *conf.Bootstrap, os *service.OxygenService) *OxygenGateway {
	return &OxygenGateway{conf: c, os: os}
}

func (og *OxygenGateway) Register(r *gin.Engine) {
	or := r.Group("/v1/oxygen", middleware.Authentication(og.conf))
	or.POST("", handler.JSON(og.os.RecordOxygen))
	or.GET("/find/:id", handler.JSON(og.os.FindOyxgen))
	or.GET("/group/:id", handler.JSON(og.os.GetOyxgens))
}
