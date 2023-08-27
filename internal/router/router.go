/*
 * @FilePath: /workflow-server/internal/router/router.go
 * @Author: maggot-code
 * @Date: 2023-08-13 02:12:24
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-22 15:46:12
 * @Description:
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/maggot-code/workflow-server/internal/middleware"
	v1 "github.com/maggot-code/workflow-server/internal/router/v1"
	"github.com/maggot-code/workflow-server/pkg/handler"
)

var ProviderSet = wire.NewSet(NewRouter, v1.NewUserGateway, v1.NewOxygenGateway, v1.NewWorkGateway)

func NewRouter(ug *v1.UserGateway, og *v1.OxygenGateway, wg *v1.WorkGateway) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middleware.RequestId())

	ug.Register(r)
	og.Register(r)
	wg.Register(r)

	r.NoRoute(handler.NotFound)
	return r
}
