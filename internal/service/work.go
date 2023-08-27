/*
 * @FilePath: /workflow-server/internal/service/work.go
 * @Author: maggot-code
 * @Date: 2023-08-20 19:13:09
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-27 16:50:26
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
	CreateWorkLossParams = errors.StatusBadRequest("create work need to 'mark_time' ", "miss params")
	CreateWorkError      = errors.StatusBadRequest("create work error", "create work error")
)

type WorkService struct {
	wc *biz.WorkUseCase
	oc *biz.OxygenUseCase
	uc *biz.UserUseCase
}

func NewWorkService(wc *biz.WorkUseCase, oc *biz.OxygenUseCase, uc *biz.UserUseCase) *WorkService {
	return &WorkService{wc: wc, oc: oc, uc: uc}
}

func (ws *WorkService) CreateWork(ctx *gin.Context, req *http.Request) (handler.Response, error) {
	var work *biz.Work

	user, err := ws.uc.NewUser(ctx)
	if err != nil {
		return nil, NewUserError.WithReason(err.Error()).WithCause(err)
	}

	if err = ctx.ShouldBindJSON(&work); err != nil {
		return nil, CreateWorkLossParams.WithCause(err)
	}

	w, err := ws.wc.CreateWork(ctx, user, work)
	if err != nil {
		return nil, CreateWorkError.WithMessage(err.Error()).WithCause(err)
	}

	if err = ws.oc.CreateOxygen(ctx, user, w); err != nil {
		return nil, CreateWorkError.WithMessage(err.Error()).WithCause(err)
	}

	return w, nil
}

func (ws *WorkService) WorkToMark(ctx *gin.Context, req *http.Request) (handler.Response, error) {
	user, err := ws.uc.NewUser(ctx)
	if err != nil {
		return nil, NewUserError.WithReason(err.Error()).WithCause(err)
	}

	mt := carbon.Parse(ctx.Param("mark_time"))

	return ws.wc.MarkTime(ctx, user, &biz.Work{MarkTime: carbon.DateTime{Carbon: mt}})
}
