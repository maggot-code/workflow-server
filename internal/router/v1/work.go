/*
 * @FilePath: /workflow-server/internal/router/v1/work.go
 * @Author: maggot-code
 * @Date: 2023-08-20 19:11:11
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-25 17:39:42
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

type WorkGateway struct {
	conf *conf.Bootstrap
	ws   *service.WorkService
}

func NewWorkGateway(c *conf.Bootstrap, ws *service.WorkService) *WorkGateway {
	return &WorkGateway{conf: c, ws: ws}
}

func (wg *WorkGateway) Register(r *gin.Engine) {
	wr := r.Group("/v1/work", middleware.Authentication(wg.conf))
	wr.PUT("", handler.JSON(wg.ws.CreateWork))
	wr.GET("/mark/:mark_time", handler.JSON(wg.ws.WorkToMark))
}
