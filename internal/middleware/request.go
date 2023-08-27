/*
 * @FilePath: /workflow-server/internal/middleware/request.go
 * @Author: maggot-code
 * @Date: 2023-08-14 16:36:22
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-15 21:18:55
 * @Description:
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := uuid.NewV4().String()
		ctx.Header("X-Request-Id", id)
		ctx.Set("X-Request-Id", id)
		ctx.Next()
	}
}
