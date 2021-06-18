package main

import (
	_ "github.com/go-playground/validator"
	"github.com/gin-gonic/gin"
	"design-api/router"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
	"design-api/config"
	"runtime"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	router.InitRouter(r)
	fmt.Println("初始化路由...")

	sysType := runtime.GOOS

	if sysType == "linux" {
		if err := r.RunUnix(config.Config.Addr.Unix); err != nil {
			fmt.Println("监听出错了:" + err.Error())
		}
	} else if sysType == "windows" {
		if err := r.Run(config.Config.Addr.Tcp); err != nil {
			fmt.Println("监听出错了:" + err.Error())
		}
	}

}
