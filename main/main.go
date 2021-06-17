package main

import (
	_ "github.com/go-playground/validator"
	"github.com/gin-gonic/gin"
	"design-api/router"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	router.InitRouter(r)
	fmt.Println("初始化路由...")
	if err := r.Run("0.0.0.0:8082"); err != nil {
		fmt.Println("监听出错了:" + err.Error())
	}
}
