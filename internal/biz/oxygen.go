/*
 * @FilePath: /workflow-server/internal/biz/oxygen.go
 * @Author: maggot-code
 * @Date: 2023-08-27 16:30:42
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-28 01:58:01
 * @Description:
 */
package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/maggot-code/workflow-server/pkg/orm"
)

type Oxygen struct {
	orm.Model
	UID        uint            `gorm:"column:uid" json:"uid"`
	WID        uint            `gorm:"column:wid" json:"wid"`
	RawData    float64         `gorm:"column:raw_data" json:"raw_data"`
	IntMap     float64         `gorm:"column:int_map" json:"int_map"`
	FloatMap   float64         `gorm:"column:float_map" json:"float_map"`
	CountData  float64         `gorm:"column:count_data" json:"count_data"`
	FixedData  float64         `gorm:"column:fixed_data" json:"fixed_data"`
	IsFixed    uint8           `gorm:"column:is_fixed" json:"is_fixed"`
	IsAutoPush uint8           `gorm:"column:is_auto_push" json:"is_auto_push"`
	State      uint8           `gorm:"column:state" json:"state"`
	EffectTime carbon.DateTime `gorm:"column:effect_time" json:"effect_time"`
}

type OxygenRecord struct {
	ID      string `json:"id" binding:"required"`
	RawData string `json:"raw_data" binding:"required"`
}

func (Oxygen) TableName() string {
	return "oxygen_records"
}

type OxygenRepo interface {
	BatchCreate(ctx *gin.Context, oxygens []*Oxygen) error
	QueryOnWork(ctx *gin.Context, w *Work) ([]*Oxygen, error)
	QueryVoid(ctx *gin.Context, o *Oxygen) ([]*Oxygen, error)
	QueryValid(ctx *gin.Context, o *Oxygen) (*Oxygen, error)
	Find(ctx *gin.Context, id uint) (*Oxygen, error)
	Save(ctx *gin.Context, o *Oxygen) error
	Record(ctx *gin.Context, o *Oxygen) error
	Finished(ctx *gin.Context, o *Oxygen) (bool, error)
}

type OxygenUseCase struct {
	repo OxygenRepo
}

func NewOxygenUseCase(or OxygenRepo) *OxygenUseCase {
	return &OxygenUseCase{repo: or}
}

func (oc *OxygenUseCase) CreateOxygen(ctx *gin.Context, u *User, w *Work) error {
	// 根据work的start_time和end_time每两小时为间隔创建一个记录并保存入库
	var oxygens []*Oxygen
	st := w.StartTime
	et := w.EndTime.AddHours(-2)

	for ct := st; ct.Compare("<=", et); {
		oxygens = append(oxygens, &Oxygen{
			UID:        u.ID,
			WID:        w.ID,
			EffectTime: ct,
		})
		ct = carbon.DateTime{Carbon: ct.AddHours(2)}
	}

	return oc.repo.BatchCreate(ctx, oxygens)
}

func (oc *OxygenUseCase) GetOxygens(ctx *gin.Context, w *Work) ([]*Oxygen, error) {
	return oc.repo.QueryOnWork(ctx, w)
}

func (oc *OxygenUseCase) FindOxygen(ctx *gin.Context, id uint) (*Oxygen, error) {
	return oc.repo.Find(ctx, id)
}

func (oc *OxygenUseCase) RecordOxygen(ctx *gin.Context, oxygen *Oxygen) error {
	// 查询无效数据
	voids, err := oc.repo.QueryVoid(ctx, oxygen)
	if err != nil {
		return err
	}

	length := len(voids)
	// 如果voids长度为0说明没有无效数据，直接保存该次上报数据
	if length <= 0 {
		return oc.repo.Record(ctx, oxygen)
	}

	// 存在无效数据，优先处理最早一条无效数据
	voidOxygen := voids[0]
	// 找到最后一条有效数据
	validOxygen, err := oc.repo.QueryValid(ctx, oxygen)
	if err != nil {
		return err
	}

	// 公式：无效原始数值 = 最后一条有效原始数值 - ( (最后一条有效原始数值 - 本次最新上报原始数值) / (无效数据条数 * 2小时间隔常量) )
	// 通过公式计算最早一条无效数据的原属数值应该是多少
	voidOxygen.RawData = validOxygen.RawData - ((validOxygen.RawData - oxygen.RawData) / float64(length*2))
	voidOxygen.IsAutoPush = 1
	if err = oc.repo.Record(ctx, voidOxygen); err != nil {
		return err
	}

	// 处理完无效数据后，保存本次上报数据
	// 递归调用RecordOxygen方法，直到无效数据全部处理完毕
	return oc.RecordOxygen(ctx, oxygen)
}

func (oc *OxygenUseCase) FinishAllRecord(ctx *gin.Context, o *Oxygen) (bool, error) {
	return oc.repo.Finished(ctx, o)
}
