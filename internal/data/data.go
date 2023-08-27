/*
 * @FilePath: /workflow-server/internal/data/data.go
 * @Author: maggot-code
 * @Date: 2023-08-13 02:16:37
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-27 16:36:30
 * @Description:
 */
package data

import (
	"github.com/google/wire"
	"github.com/maggot-code/workflow-server/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewDB, NewUserRepo, NewWorkRepo, NewOxygenReop)

type Data struct {
	db *gorm.DB
}

func NewData(c *conf.Bootstrap, db *gorm.DB) (*Data, func(), error) {
	return &Data{db: db}, func() {}, nil
}

func NewDB(c *conf.Bootstrap) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Data.Database.Source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
		PrepareStmt:                              true,
	})
	if err != nil {
		panic(err)
	}

	return db
}
