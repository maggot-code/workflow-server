/*
 * @FilePath: /workflow-server/internal/data/oxygen.go
 * @Author: maggot-code
 * @Date: 2023-08-27 16:35:27
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-28 01:11:39
 * @Description:
 */
package data

import (
	"github.com/gin-gonic/gin"
	"github.com/maggot-code/workflow-server/internal/biz"
)

var _ biz.OxygenRepo = (*oxygenRepo)(nil)

type oxygenRepo struct {
	data *Data
}

func NewOxygenReop(data *Data) biz.OxygenRepo {
	return &oxygenRepo{data: data}
}

func (or *oxygenRepo) BatchCreate(ctx *gin.Context, oxygens []*biz.Oxygen) error {
	return or.data.db.CreateInBatches(oxygens, len(oxygens)).Error
}

func (or *oxygenRepo) QueryOnWork(ctx *gin.Context, w *biz.Work) ([]*biz.Oxygen, error) {
	var oxygens []*biz.Oxygen

	if err := or.data.db.
		Model(&biz.Oxygen{}).
		Where("uid = ? AND wid = ?", w.UID, w.ID).
		Order("effect_time ASC").
		Find(&oxygens).Error; err != nil {
		return nil, err
	}

	return oxygens, nil
}

// 查询有效数据
func (or *oxygenRepo) QueryValid(ctx *gin.Context, o *biz.Oxygen) (*biz.Oxygen, error) {
	var oxygens *biz.Oxygen

	if err := or.data.db.
		Model(&biz.Oxygen{}).
		Where("state = ? AND wid = ? AND effect_time < ?", 1, o.WID, o.EffectTime).
		Order("effect_time DESC").
		First(&oxygens).Error; err != nil {
		return nil, err
	}

	return oxygens, nil
}

// 查询无效数据
func (or *oxygenRepo) QueryVoid(ctx *gin.Context, o *biz.Oxygen) ([]*biz.Oxygen, error) {
	var oxygens []*biz.Oxygen

	if err := or.data.db.
		Model(&biz.Oxygen{}).
		Where("state = ? AND wid = ? AND effect_time < ?", 0, o.WID, o.EffectTime).
		Order("effect_time ASC").
		Find(&oxygens).Error; err != nil {
		return nil, err
	}

	return oxygens, nil
}

func (or *oxygenRepo) Find(ctx *gin.Context, id uint) (*biz.Oxygen, error) {
	var oxygen biz.Oxygen

	if err := or.data.db.
		Model(&biz.Oxygen{}).
		Where("id = ?", id).
		First(&oxygen).Error; err != nil {
		return nil, err
	}

	return &oxygen, nil
}

func (or *oxygenRepo) Save(ctx *gin.Context, o *biz.Oxygen) error {
	return or.data.db.Save(o).Error
}

func (or *oxygenRepo) Record(ctx *gin.Context, oxygen *biz.Oxygen) error {
	// 保存该次数据并返回响应
	intval, err := RecordToInt(oxygen.RawData)
	if err != nil {
		return err
	}

	floatval, err := RecordToFloat(oxygen.RawData)
	if err != nil {
		return err
	}

	oxygen.State = 1
	oxygen.IntMap = intval
	oxygen.FloatMap = floatval
	oxygen.CountData = intval + floatval

	return or.Save(ctx, oxygen)
}

func (or *oxygenRepo) Finished(ctx *gin.Context, o *biz.Oxygen) (bool, error) {
	var count int64

	if err := or.data.db.Model(&biz.Oxygen{}).Where("wid = ? AND state = ?", o.WID, 0).Count(&count); err.Error != nil {
		return false, err.Error
	}

	return count <= 0, nil
}
