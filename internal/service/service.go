/*
 * @FilePath: /workflow-server/internal/service/service.go
 * @Author: maggot-code
 * @Date: 2023-08-13 02:14:29
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-20 21:48:27
 * @Description:
 */
package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserService, NewWorkService, NewOxygenService)
