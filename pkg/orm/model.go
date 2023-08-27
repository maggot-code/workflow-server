/*
 * @FilePath: /workflow-server/pkg/orm/model.go
 * @Author: maggot-code
 * @Date: 2023-08-14 23:18:20
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-20 22:09:00
 * @Description:
 */
package orm

import (
	"github.com/golang-module/carbon/v2"
)

type Model struct {
	ID        uint            `gorm:"primarykey" json:"id"`
	CreatedAt carbon.DateTime `gorm:"column:created_at" json:"created_at"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt carbon.DateTime `gorm:"column:deleted_at;default:NULL" json:"-"`
}
