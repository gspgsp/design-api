package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

/**
解析json参数
*/
func JsonParamParse(c *gin.Context) map[string]string {
	param, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(param, &m)

	return m
}
