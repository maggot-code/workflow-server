/*
 * @FilePath: /workflow-server/internal/biz/work.go
 * @Author: maggot-code
 * @Date: 2023-08-20 19:14:47
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-28 01:18:09
 * @Description:
 */
package biz

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/maggot-code/workflow-server/pkg/orm"
)

type Work struct {
	orm.Model
	UID       uint            `gorm:"column:uid" json:"uid"`
	State     uint8           `gorm:"column:state" json:"state"`
	StartTime carbon.DateTime `gorm:"column:start_time" json:"start_time"`
	EndTime   carbon.DateTime `gorm:"column:end_time" json:"end_time"`
	MarkTime  carbon.DateTime `gorm:"column:mark_time" json:"mark_time" binding:"required" time_format:"2006-01-02 15:04:05"`
}

func (Work) TableName() string {
	return "weekdays"
}

type WorkRepo interface {
	Exist(ctx *gin.Context, u *User, w *Work) (bool, error)
	Save(ctx *gin.Context, w *Work) (*Work, error)
	Find(ctx *gin.Context, id uint) (*Work, error)
	FindToMark(ctx *gin.Context, u *User, w *Work) (*Work, error)
}

type WorkUseCase struct {
	repo WorkRepo
}

func NewWorkUseCase(wr WorkRepo) *WorkUseCase {
	return &WorkUseCase{repo: wr}
}

func (wc *WorkUseCase) CreateWork(ctx *gin.Context, u *User, w *Work) (*Work, error) {
	exist, err := wc.repo.Exist(ctx, u, w)
	if err != nil {
		return nil, err
	}

	if exist {
		return nil, fmt.Errorf("work is exist")
	}

	w.MarkTime = carbon.DateTime{Carbon: w.MarkTime.StartOfDay()}
	w.StartTime = carbon.DateTime{Carbon: w.MarkTime.StartOfDay().AddHours(8)}
	w.EndTime = carbon.DateTime{Carbon: w.StartTime.AddDay()}
	w.UID = u.ID

	return wc.repo.Save(ctx, w)
}

func (wc *WorkUseCase) MarkTime(ctx *gin.Context, u *User, w *Work) (*Work, error) {
	return wc.repo.FindToMark(ctx, u, w)
}

func (wc *WorkUseCase) GetWork(ctx *gin.Context, id uint) (*Work, error) {
	return wc.repo.Find(ctx, id)
}

func (wc *WorkUseCase) UpdateState(ctx *gin.Context, w *Work) error {
	w.State = 1

	_, err := wc.repo.Save(ctx, w)

	return err
}
