package main

import (
	_ "design-api/common/log"
	"design-api/config"
	"design-api/router"
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
	"syscall"
)

var sysType string

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	router.InitRouter(r)
	log.Println("初始化路由...")

	sysType = runtime.GOOS

	if sysType == "linux" {
		syscall.Unlink(config.Config.Addr.Unix)
		err := r.RunUnix(config.Config.Addr.Unix)
		if err != nil {
			log.Println("监听出错了:" + err.Error())
		} else {
			syscall.Chmod(config.Config.Addr.Unix, 0777)
		}
	} else if sysType == "windows" {
		if err := r.Run(config.Config.Addr.Tcp); err != nil {
			log.Println("监听出错了:" + err.Error())
		}
	}

}
