/*
 * @FilePath: /workflow-server/internal/data/work.go
 * @Author: maggot-code
 * @Date: 2023-08-20 19:26:22
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-27 17:51:53
 * @Description:
 */
package data

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maggot-code/workflow-server/internal/biz"
)

var _ biz.WorkRepo = (*workRepo)(nil)

type workRepo struct {
	data *Data
}

func NewWorkRepo(data *Data) biz.WorkRepo {
	return &workRepo{data: data}
}

func (wr *workRepo) Exist(ctx *gin.Context, u *biz.User, w *biz.Work) (bool, error) {
	var count int64

	if err := wr.data.db.Model(&biz.Work{}).Where("uid = ? AND DATE(mark_time) = DATE(?)", u.ID, w.MarkTime).Count(&count); err.Error != nil {
		return false, fmt.Errorf("db: exist work fail; %w", err.Error)
	}

	return count > 0, nil
}

func (wr *workRepo) Save(ctx *gin.Context, w *biz.Work) (*biz.Work, error) {
	if err := wr.data.db.Save(w).Error; err != nil {
		return nil, fmt.Errorf("db: save work fail; %w", err)
	}

	return w, nil
}

func (wr *workRepo) FindToMark(ctx *gin.Context, u *biz.User, w *biz.Work) (*biz.Work, error) {
	var work biz.Work

	if err := wr.data.db.Model(&biz.Work{}).Where("uid = ? AND DATE(mark_time) = DATE(?)", u.ID, w.MarkTime).First(&work).Error; err != nil {
		return nil, fmt.Errorf("db: find work fail; %w", err)
	}

	return &work, nil
}

func (wr *workRepo) Find(ctx *gin.Context, id uint) (*biz.Work, error) {
	var work biz.Work

	if err := wr.data.db.Model(&biz.Work{}).Where("id = ?", id).First(&work).Error; err != nil {
		return nil, fmt.Errorf("db: find work fail; %w", err)
	}

	return &work, nil
}
