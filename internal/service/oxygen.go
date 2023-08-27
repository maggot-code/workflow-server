/*
 * @FilePath: /workflow-server/internal/service/oxygen.go
 * @Author: maggot-code
 * @Date: 2023-08-20 21:45:48
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-28 01:19:40
 * @Description:
 */
package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maggot-code/workflow-server/internal/biz"
	"github.com/maggot-code/workflow-server/pkg/errors"
	"github.com/maggot-code/workflow-server/pkg/handler"
)

var (
	OxygenError = errors.StatusBadRequest("oxygen error", "oxygen error")
)

type OxygenService struct {
	oc *biz.OxygenUseCase
	wc *biz.WorkUseCase
}

func NewOxygenService(oc *biz.OxygenUseCase, wc *biz.WorkUseCase) *OxygenService {
	return &OxygenService{oc: oc, wc: wc}
}

func (os *OxygenService) GetOyxgens(ctx *gin.Context, req *http.Request) (handler.Response, error) {
	wid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil, OxygenError.WithReason(err.Error()).WithCause(err)
	}

	work, err := os.wc.GetWork(ctx, uint(wid))
	if err != nil {
		return nil, OxygenError.WithReason(err.Error()).WithCause(err)
	}

	return os.oc.GetOxygens(ctx, work)
}

func (os *OxygenService) FindOyxgen(ctx *gin.Context, req *http.Request) (handler.Response, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil, OxygenError.WithReason(err.Error()).WithCause(err)
	}

	return os.oc.FindOxygen(ctx, uint(id))
}

func (os *OxygenService) RecordOxygen(ctx *gin.Context, req *http.Request) (handler.Response, error) {
	var or *biz.OxygenRecord

	if err := ctx.ShouldBindJSON(&or); err != nil {
		return nil, OxygenError.WithReason(err.Error()).WithCause(err)
	}

	id, err := strconv.Atoi(or.ID)
	if err != nil {
		return nil, OxygenError.WithReason(err.Error()).WithCause(err)
	}

	value, err := strconv.ParseFloat(or.RawData, 64)
	if err != nil {
		return nil, OxygenError.WithReason(err.Error()).WithCause(err)
	}

	oxygen, err := os.oc.FindOxygen(ctx, uint(id))
	if err != nil {
		return nil, OxygenError.WithReason(err.Error()).WithCause(err)
	}

	work, err := os.wc.GetWork(ctx, oxygen.WID)
	if err != nil {
		return nil, OxygenError.WithReason(err.Error()).WithCause(err)
	}

	oxygen.RawData = value
	if err = os.oc.RecordOxygen(ctx, oxygen); err != nil {
		return nil, OxygenError.WithReason(err.Error()).WithCause(err)
	}

	finish, err := os.oc.FinishAllRecord(ctx, oxygen)
	if err != nil {
		return nil, OxygenError.WithReason(err.Error()).WithCause(err)
	}

	if finish {
		return nil, os.wc.UpdateState(ctx, work)
	}

	return nil, nil
}
