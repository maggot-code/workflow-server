/*
 * @FilePath: /workflow-server/internal/biz/biz.go
 * @Author: maggot-code
 * @Date: 2023-08-13 02:15:06
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-27 16:36:51
 * @Description:
 */
package biz

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserUseCase, NewOxygenUseCase, NewWorkUseCase)
