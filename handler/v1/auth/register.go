package auth

import (
	"github.com/gin-gonic/gin"
	"design-api/util"
	"log"
)

func Register(c *gin.Context) {

	token, err := util.GenerateToken("1","guo")

	if err != nil {
		log.Println("token 生成错误:" +err.Error())

		c.JSON(419, gin.H{
			"error": err.Error(),
		})
	}else {
		c.JSON(200, gin.H{
			"token": token,
		})
	}
}
