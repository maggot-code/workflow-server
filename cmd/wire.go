//go:build wireinject
// +build wireinject

/*
 * @FilePath: /workflow-server/cmd/wire.go
 * @Author: maggot-code
 * @Date: 2023-08-13 02:10:50
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-14 15:18:49
 * @Description:
 */

package main

import (
	"github.com/maggot-code/workflow-server/internal/biz"
	"github.com/maggot-code/workflow-server/internal/conf"
	"github.com/maggot-code/workflow-server/internal/data"
	"github.com/maggot-code/workflow-server/internal/router"
	"github.com/maggot-code/workflow-server/internal/service"

	"github.com/google/wire"
)

func wireApp(*conf.Bootstrap) (*App, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		router.ProviderSet,
		newApp,
	))
}
