/*
 * @FilePath: /workflow-server/internal/middleware/auth.go
 * @Author: maggot-code
 * @Date: 2023-08-15 21:17:44
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-22 15:43:35
 * @Description:
 */
package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maggot-code/workflow-server/internal/conf"
	"github.com/maggot-code/workflow-server/pkg/handler"
	"github.com/maggot-code/workflow-server/pkg/jwt"
)

func Authentication(conf *conf.Bootstrap) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if len(auth) == 0 {
			handler.Unauthorized(ctx)
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(conf.Wechat.Secret, auth)
		if err != nil {
			handler.Unauthorized(ctx)
			ctx.Abort()
			return
		}

		if jwt.NeedRefresh(20*time.Minute, token) {
			token, err := jwt.Refresh(conf.Wechat.Secret, time.Now().Add(2*time.Hour), token)
			if err != nil {
				fmt.Printf("middleware: refresh token error; %v", err.Error())
				return
			}

			ctx.Header("Set-Authorization", token)
			fmt.Printf("middleware: refresh token; %v", token)
		}

		ctx.Set("metadata", token.Metadata)
		ctx.Next()
	}
}
