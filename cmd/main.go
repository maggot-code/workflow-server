/*
 * @FilePath: /workflow-server/cmd/main.go
 * @Author: maggot-code
 * @Date: 2023-08-13 02:10:01
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-28 22:30:38
 * @Description:
 */
package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maggot-code/workflow-server/internal/conf"
)

var (
	Name     = "adams.app"
	flagconf string
)

type App struct {
	router *gin.Engine
}

func init() {
	flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
}

func newApp(r *gin.Engine) *App {
	return &App{
		router: r,
	}
}

func main() {
	flag.Parse()

	conf, err := conf.New(flagconf)
	if err != nil {
		panic(err)
	}

	fmt.Println(conf)
	app, cleanup, err := wireApp(conf)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.router.Run(conf.Server.Addr); err != nil {
		panic(err)
	}
}
