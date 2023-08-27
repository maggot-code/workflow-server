/*
 * @FilePath: /workflow-server/internal/conf/conf.go
 * @Author: maggot-code
 * @Date: 2023-08-13 02:06:14
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-13 02:06:16
 * @Description:
 */
package conf

import (
	"github.com/spf13/viper"
)

func New(flagconf string) (*Bootstrap, error) {
	v := viper.New()

	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath(flagconf)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Bootstrap
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
