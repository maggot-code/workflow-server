/*
 * @FilePath: /workflow-server/pkg/handler/handler.go
 * @Author: maggot-code
 * @Date: 2023-08-14 16:03:25
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-16 18:22:54
 * @Description:
 */
package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maggot-code/workflow-server/pkg/errors"
)

var (
	NotFoundError     = errors.StatusNotFound("Not found", "not found")
	UnauthorizedError = errors.StatusUnauthorized("Unauthorized", "unauthorized")
	UnknownError      = errors.StatusInternalServerError("Unknown error", "server error")
)

type Response interface{}

type reply struct {
	errors.Status
	Data interface{} `json:"data"`
}

type handler func(ctx *gin.Context, req *http.Request) (Response, error)

func NotFound(ctx *gin.Context) {
	ctx.JSON(int(NotFoundError.Code), reply{
		Status: errors.Status{
			Code:    NotFoundError.Code,
			Message: NotFoundError.Message,
		},
	})
}

func Unauthorized(ctx *gin.Context) {
	ctx.JSON(int(UnauthorizedError.Code), reply{
		Status: errors.Status{
			Code:    UnauthorizedError.Code,
			Message: UnauthorizedError.Message,
		},
	})
}

func JSON(hd handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := hd(ctx, ctx.Request)
		we := errors.FromError(err)

		if we != nil {
			fmt.Println(we.Error())
			fmt.Println(we.Unwrap())
			ctx.JSON(int(we.Code), reply{
				Status: errors.Status{
					Code:    we.Code,
					Message: we.Message,
				},
				Data: data,
			})
			return
		}

		ok := errors.StatusOK("", "ok")
		ctx.JSON(int(ok.Code), reply{
			Status: errors.Status{
				Code:    ok.Code,
				Message: ok.Message,
			},
			Data: data,
		})
	}
}
