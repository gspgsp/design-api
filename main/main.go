package main

import (
	"github.com/gin-gonic/gin"
	"design-api/router"
	"design-api/config"
	"runtime"
	"log"
	_ "design-api/common/log"
)

var sysType string

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	router.InitRouter(r)
	log.Println("初始化路由...")

	sysType = runtime.GOOS

	if sysType == "linux" {
		if err := r.RunUnix(config.Config.Addr.Unix); err != nil {
			log.Println("监听出错了:" + err.Error())
		}
	} else if sysType == "windows" {
		if err := r.Run(config.Config.Addr.Tcp); err != nil {
			log.Println("监听出错了:" + err.Error())
		}
	}

}
